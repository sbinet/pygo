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
		code:     []byte{124, 0, 0, 124, 1, 0, 23, 83},
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
	}

	for _, test := range []struct {
		name string
		x    Value
		y    Value
		want Value
	}{
		{
			name: "add-bb",
			x:    false,
			y:    false,
			want: 0,
		},
		{
			name: "add-bb",
			x:    true,
			y:    false,
			want: 1,
		},
		{
			name: "add-bb",
			x:    true,
			y:    true,
			want: 2,
		},
		{
			name: "add-xx",
			x:    byte(104),
			y:    byte(111),
			want: byte(215),
		},
		{
			name: "add-ii",
			x:    -44,
			y:    +2,
			want: -42,
		},
		{
			name: "add-ii",
			x:    44,
			y:    -2,
			want: 42,
		},
		{
			name: "add-ii",
			x:    40,
			y:    2,
			want: 42,
		},
		{
			name: "add-fi",
			x:    40,
			y:    2.0,
			want: 42.0,
		},
		{
			name: "add-ff",
			x:    40.1,
			y:    2.0,
			want: 42.1,
		},
		{
			name: "add-ss",
			x:    "hel",
			y:    "lo",
			want: "hello",
		},
	} {
		locals["x"] = test.x
		locals["y"] = test.y
		v, err := vm.RunCode(code, globals, locals)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(v, test.want) {
			t.Fatalf("%s(%v, %v): got %v. want %v\n", test.name, test.x, test.y, v, test.want)
		}
	}
}
