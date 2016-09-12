// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen-opcodes.go
//go:generate goimports -w opcodes_gen.go

// pygo exposes an API to interpret Python bytecode.
package pygo

// VM is the pygo virtual machine implementation.
type VM struct {
	frames Frames    // call stack of Frames
	fp     *Frame    // pointer to current Frame
	ret    Value     // return value
	exc    exception // last exception
}

type Frame struct {
	code     Code
	globals  map[string]Value
	locals   map[string]Value
	builtins map[string]Value
	prev     *Frame
	ip       int // instruction pointer

	stack  []Value
	blocks []block
}

// Frames is a stack of frames
type Frames []Frame

type Value interface{}

type exception interface {
	Type() Value
	Value() Value
	Traceback() Value
}

type pyException struct {
	typ Value
	val Value
	tb  Value
}

type block struct {
	typ  blockKind
	next int
	stk  int
}

type blockKind byte

const (
	bkLoop blockKind = iota
	bkExceptHandler
	bkSetupExcept
	bkFinally
	bkWith
)

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

// Opcode is a single bytecode operation for the Python interpreter.
// Operands (if any) for the opcode follow in the bytecode stream.
type Opcode byte

// Instruction is a single instruction in a bytecode stream.
type Instruction interface{}

// New creates a new Python VM.
func New() *VM {
	var vm VM
	return &vm
}

// RunCode interprets code compiled to bytecode by some means.
func (vm *VM) RunCode(code Code) (Value, error) {
	return nil, nil
}

type stack struct {
	stk []Value
}

func (s *stack) push(v Value) {
	s.stk = append(s.stk, v)
}

func (s *stack) pop() Value {
	i := len(s.stk) - 1
	v := s.stk[i]
	s.stk = s.stk[:i]
	return v
}
