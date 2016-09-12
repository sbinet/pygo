// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

type Frame struct {
	code     *Code
	globals  map[string]Value
	locals   map[string]Value
	builtins map[string]Value
	prev     *Frame
	ip       int // instruction pointer

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
	// FIXME(sbinet): handle cells
	// FIXME(sbinet): handle generator
	return frame
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
