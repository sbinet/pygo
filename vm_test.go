// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

import (
	"reflect"
	"testing"
)

func TestVM(t *testing.T) {
	var vm VM
	code := &Code{
		instr:    []byte{124, 0, 0, 124, 1, 0, 23, 83},
		varnames: []string{"x", "y"},
	}
	globals := map[string]Value{
		"__builtins__": builtins,
		"__name__":     "__main__",
		"__doc__":      nil,
		"__package__":  nil,
	}
	locals := map[string]Value{
		"__builtins__": builtins,
		"x":            40,
		"y":            2,
	}
	v, err := vm.RunCode(code, globals, locals)
	if err != nil {
		t.Fatal(err)
	}
	want := 42
	if !reflect.DeepEqual(v, want) {
		t.Fatalf("got %v. want %v\n", v, want)
	}
}
