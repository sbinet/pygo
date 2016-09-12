// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen-opcodes.go
//go:generate goimports -w opcodes_gen.go

// pygo exposes an API to interpret Python bytecode.
package pygo

type Value interface{}

type exception struct {
	typ Value
	val Value
	tb  Value
}

var (
	// FIXME(sbinet): populate builtins
	builtins = make(map[string]Value)
)
