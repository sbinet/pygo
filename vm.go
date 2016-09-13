// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

import "fmt"

// Code represents byte-compiled executable Python code.
type Code struct {
	name     string   // function name
	code     []byte   // bytecode instructions
	nargs    int      // number of positional arguments (including arguments with default values)
	nlocals  int      // number of local variables used by the function (including arguments)
	nkwargs  int      // number of keyword arguments
	consts   []Value  // tuple containing the literals used by the bytecode
	names    []string // tuple containing the names used by the bytecode
	varnames []string // tuple containing the names of the local variables (starting with the argument names)
	cellvars []string // tuple containing the names of local variables that are referenced by nested functions
	freevars []string // tuple containing the names of free variables

	fname  string // filename from which the code was compiled
	lineno int    // first line number of the function
	flags  int    // number of flags for the interpreter

	lnotab  []byte // slice encoding the mapping from bytecode offsets to line numbers
	stacksz int    // stacksize is the required stack size (including local variables)
}

type stack []Value

func (s *stack) push(v Value) {
	*s = append(*s, v)
}

func (s *stack) pushn(v ...Value) {
	*s = append(*s, v...)
}

func (s *stack) pop() Value {
	i := len(*s) - 1
	v := (*s)[i]
	*s = (*s)[:i]
	return v
}

func (s *stack) popn(n int) []Value {
	o := make([]Value, n)
	i := len(*s) - n
	copy(o, (*s)[i:])
	*s = (*s)[:i]
	return o
}

func (s stack) top() Value {
	return s[len(s)-1]
}

// Opcode is a single bytecode operation for the Python interpreter.
// Operands (if any) for the opcode follow in the bytecode stream.
type Opcode byte

// VM is the pygo virtual machine implementation.
type VM struct {
	frames Frames     // call stack of Frames
	fp     *Frame     // pointer to current Frame
	ret    Value      // return value
	exc    *exception // last exception
}

// New creates a new Python VM.
func New() *VM {
	return &VM{}
}

// RunCode interprets code compiled to bytecode by some means.
func (vm *VM) RunCode(code *Code, globals, locals map[string]Value) (Value, error) {
	frame := newFrame(code, nil, globals, locals, vm.fp)
	val, err := vm.runFrame(frame)
	// FIXME(sbinet): check frames leftovers
	// FIXME(sbinet): check data left on stack
	return val, err
}

func (vm *VM) pushFrame(f *Frame) {
	vm.frames.push(f)
	vm.fp = f
}

func (vm *VM) popFrame() {
	vm.frames.pop()
	if len(vm.frames) > 0 {
		vm.fp = vm.frames[len(vm.frames)-1]
	} else {
		vm.fp = nil
	}
}

func (vm *VM) resumeFrame(f *Frame) (Value, error) {
	f.prev = vm.fp
	val, err := vm.runFrame(f)
	f.prev = nil
	return val, err
}

func (vm *VM) runFrame(f *Frame) (Value, error) {
	vm.pushFrame(f)
	defer vm.popFrame()

	for f.ip < len(f.code.code) {
		op := Opcode(f.code.code[f.ip])
		f.ip++

		switch op {
		case Op_LOAD_CONST:
			i := f.iload()
			v := f.code.consts[i]
			f.stack.push(v)

		case Op_POP_TOP:
			_ = f.stack.pop()

		case Op_DUP_TOP:
			v := f.stack.top()
			f.stack.push(v)

		case Op_ROT_TWO:
			v := f.stack.popn(2)
			f.stack.pushn(v[1], v[0])

		case Op_ROT_THREE:
			v := f.stack.popn(3)
			f.stack.pushn(v[2], v[0], v[1])

		case Op_LOAD_NAME:
			i := f.iload()
			n := f.code.names[i]
			v, err := f.loadName(n)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_STORE_NAME:
			i := f.iload()
			n := f.code.names[i]
			v := f.stack.pop()
			f.locals[n] = v

		case Op_DELETE_NAME:
			i := f.iload()
			n := f.code.names[i]
			delete(f.locals, n)

		case Op_LOAD_FAST:
			i := f.iload()
			n := f.code.varnames[i]
			v, ok := f.locals[n]
			if !ok {
				// FIXME(sbinet): better error
				return nil, fmt.Errorf(
					"local variable '%s' referenced before assignment",
					n,
				)
			}
			f.stack.push(v)

		case Op_STORE_FAST:
			i := f.iload()
			n := f.code.varnames[i]
			f.locals[n] = f.stack.pop()

		case Op_DELETE_FAST:
			i := f.iload()
			n := f.code.varnames[i]
			delete(f.locals, n)

		case Op_LOAD_GLOBAL:
			i := f.iload()
			n := f.code.names[i]
			v, err := f.loadGlobal(n)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_LOAD_DEREF:
			i := f.iload()
			var n string
			if i < len(f.code.cellvars) {
				n = f.code.cellvars[i]
			} else {
				i -= len(f.code.cellvars)
				n = f.code.freevars[i]
			}
			v := f.cells[n]
			f.stack.push(v)

		case Op_STORE_DEREF:
			i := f.iload()
			var n string
			if i < len(f.code.cellvars) {
				n = f.code.cellvars[i]
			} else {
				i -= len(f.code.cellvars)
				n = f.code.freevars[i]
			}
			f.cells[n] = f.stack.pop()

		case Op_UNARY_INVERT:
			v := f.stack.pop()
			v, err := unaryInvert(v)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_UNARY_NEGATIVE:
			v := f.stack.pop()
			v, err := unaryNegative(v)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_UNARY_NOT:
			v := f.stack.pop()
			v, err := unaryNot(v)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_UNARY_POSITIVE:
			v := f.stack.pop()
			v, err := unaryPositive(v)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_BINARY_ADD, Op_BINARY_AND, Op_BINARY_FLOOR_DIVIDE, Op_BINARY_LSHIFT,
			Op_BINARY_MATRIX_MULTIPLY, Op_BINARY_MODULO, Op_BINARY_MULTIPLY,
			Op_BINARY_OR, Op_BINARY_POWER, Op_BINARY_RSHIFT, Op_BINARY_SUBSCR,
			Op_BINARY_SUBTRACT,
			Op_BINARY_TRUE_DIVIDE, Op_BINARY_XOR:
			args := f.stack.popn(2)
			v, err := binaryOp(args[0], args[1], op)
			if err != nil {
				return nil, err
			}
			f.stack.push(v)

		case Op_RETURN_VALUE:
			vm.ret = f.stack.pop()

		default:
			return nil, fmt.Errorf("pygo: internal error: unknown bytecode %[1]d (%[1]v)", byte(op), op)
		}
	}

	return vm.ret, nil
}
