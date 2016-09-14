// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

// Object defines the PyObject protocol.
type Object interface {
	HasAttr(n string) bool
	GetAttr(n string) (Object, error)
	SetAttr(n string, v Object) error
	DelAttr(n string) error

	Repr() string
	Str() string

	Call(args, kwds Object) (Object, error)

	Length() int

	GetItem(key Object) (Object, error)
	SetItem(key, value Object) error
	DelItem(key Object) error

	Dir() ([]string, error)
}
