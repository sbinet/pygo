// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

// Code represents byte-compiled executable Python code.
type Code struct {
	name     string   // function name
	instr    []byte   // bytecode instructions
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
	for {
	}
	vm.popFrame()

	return vm.ret, nil
}
