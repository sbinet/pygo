// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

import (
	"encoding/binary"
	"fmt"
)

type Frame struct {
	code     *Code
	globals  map[string]Value
	locals   map[string]Value
	builtins map[string]Value
	prev     *Frame
	ip       int // instruction pointer

	cells  map[string]Value
	stack  stack
	blocks []block
}

func newFrame(code *Code, callargs, globals, locals map[string]Value, prev *Frame) *Frame {
	switch {
	case globals != nil:
		if locals == nil {
			locals = globals
		}
	case prev != nil:
		globals = prev.globals
		locals = make(map[string]Value)
	default:
		globals = map[string]Value{
			"__builtins__": builtins,
			"__name__":     "__main__",
			"__doc__":      nil,
			"__package__":  nil,
		}
		locals = globals
	}
	for k, v := range callargs {
		locals[k] = v
	}

	frame := &Frame{
		code:    code,
		globals: globals,
		locals:  locals,
		prev:    prev,
	}

	if prev != nil {
		frame.builtins = prev.builtins
	} else {
		frame.builtins = frame.locals["__builtins__"].(map[string]Value)
	}

	if len(frame.code.cellvars) > 0 {
		frame.cells = make(map[string]Value)
		if frame.prev.cells == nil {
			frame.prev.cells = make(map[string]Value)
		}
		for _, n := range frame.code.cellvars {
			v := frame.locals[n]
			frame.cells[n] = v
			frame.prev.cells[n] = v
		}
	}

	if len(frame.code.freevars) > 0 {
		if frame.cells == nil {
			frame.cells = make(map[string]Value)
		}
		for _, n := range frame.code.freevars {
			frame.cells[n] = frame.prev.cells[n]
		}
	}

	// FIXME(sbinet): handle generator
	return frame
}

func (f *Frame) iload() int {
	i := int(binary.LittleEndian.Uint16(f.code.code[f.ip : f.ip+2]))
	f.ip += 2
	return i
}

func (f *Frame) loadName(n string) (Value, error) {
	v, ok := f.locals[n]
	if ok {
		return v, nil
	}
	v, ok = f.globals[n]
	if ok {
		return v, nil
	}
	v, ok = f.builtins[n]
	if ok {
		return n, nil
	}
	return nil, fmt.Errorf("name '%s' is not defined", n)
}

func (f *Frame) loadGlobal(n string) (Value, error) {
	v, ok := f.globals[n]
	if ok {
		return v, nil
	}
	v, ok = f.builtins[n]
	if ok {
		return n, nil
	}
	return nil, fmt.Errorf("global name '%s' is not defined", n)
}

// Frames is a stack of frames
type Frames []*Frame

func (fs *Frames) push(f *Frame) {
	*fs = append(*fs, f)
}

func (fs *Frames) pop() *Frame {
	i := len(*fs) - 1
	v := (*fs)[i]
	*fs = (*fs)[:i]
	return v
}

type block struct {
	typ  blockKind
	next int
	stk  int
}

type blockKind byte

const (
	bkInvalid blockKind = iota
	bkLoop
	bkExceptHandler
	bkSetupExcept
	bkFinally
	bkWith
)
