// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	// code: the set of instructions to execute '7+5'
	code := Code{
		Prog: []Instruction{
			OpLoadValue, 0, // load first number
			OpLoadValue, 1, // load second number
			OpAdd,
			OpPrint,
		},
		Numbers: []int{7, 5},
	}

	var interp Interpreter
	interp.Run(code)
}

// Opcode is a single bytecode operation for the tiny interpreter.
// Operands (if any) for the opcode follow in the bytecode stream.
type Opcode int

const (
	OpLoadValue Opcode = iota
	OpAdd
	OpPrint
)

// Instruction is a single instruction in a bytecode stream.
type Instruction interface{}

// Code is a complete bytecompiled program together with data.
type Code struct {
	Prog    []Instruction // Prog is the set of instructions to execute.
	Numbers []int         // Numbers is the data being manipulated by the program.
}

// Interpreter interprets instructions for the tiny-interp interpreter.
type Interpreter struct {
	stack stack
}

func (interp *Interpreter) Run(code Code) {
	prog := code.Prog
	for pc := 0; pc < len(prog); pc++ {
		op := prog[pc].(Opcode)
		switch op {
		case OpLoadValue:
			pc++
			val := code.Numbers[prog[pc].(int)]
			interp.stack.push(val)
		case OpAdd:
			lhs := interp.stack.pop()
			rhs := interp.stack.pop()
			sum := lhs + rhs
			interp.stack.push(sum)
		case OpPrint:
			val := interp.stack.pop()
			fmt.Println(val)
		}
	}
}

type stack struct {
	stk []int
}

func (s *stack) push(v int) {
	s.stk = append(s.stk, v)
}

func (s *stack) pop() int {
	i := len(s.stk) - 1
	v := s.stk[i]
	s.stk = s.stk[:i]
	return v
}
