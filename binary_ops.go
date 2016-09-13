// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Automatically generated with 'go generate'.
// DO NOT MODIFY

package pygo

import (
	"fmt"
	"reflect"
)

func binaryOp(a, b Value, op Opcode) (Value, error) {
	rva := reflect.ValueOf(a)
	rvb := reflect.ValueOf(b)
	rva, rvb, err := promote(rva, rvb)
	if err != nil {
		return nil, err
	}

	switch rva.Interface().(type) {
	case bool:
		aa := bool2int(rva.Bool())
		bb := bool2int(rvb.Bool())
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}

	case byte:
		aa := byte(rva.Uint())
		bb := byte(rvb.Uint())
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}
	case int:
		aa := int(rva.Int())
		bb := int(rvb.Int())
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}
	case int64:
		aa := rva.Int()
		bb := rvb.Int()
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}

	case float32:
		aa := float32(rva.Float())
		bb := float32(rvb.Float())
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}
	case float64:
		aa := rva.Float()
		bb := rvb.Float()
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}

	case complex64:
		aa := complex64(rva.Complex())
		bb := complex64(rvb.Complex())
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}
	case complex128:
		aa := rva.Complex()
		bb := rvb.Complex()
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}

	case string:
		aa := rva.String()
		bb := rvb.String()
		switch op {
		case Op_BINARY_ADD:
			return aa + bb, nil
		default:
			panic("not implemented")
		}

	}

	panic("not implemented")
}

func promote(a, b reflect.Value) (reflect.Value, reflect.Value, error) {
	ak := a.Kind()
	bk := b.Kind()
	if ak == bk {
		return a, b, nil
	}

	cnv, ok := pmap[pmapKey{ak, bk}]
	if ok {
		a, b = cnv(a, b)
		return a, b, nil
	}

	cnv, ok = pmap[pmapKey{bk, ak}]
	if ok {
		b, a = cnv(b, a)
		return a, b, nil
	}

	return reflect.Value{}, reflect.Value{}, fmt.Errorf("pygo: no valid python type promotion (%v, %v)", ak, bk)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

type pmapKey struct {
	a reflect.Kind
	b reflect.Kind
}

type pmapValue func(a, b reflect.Value) (reflect.Value, reflect.Value)

type pmapType map[pmapKey]pmapValue

var pmap pmapType

func init() {
	pmap = map[pmapKey]pmapValue{
		{reflect.Bool, reflect.Int}: func(a, b reflect.Value) (reflect.Value, reflect.Value) {
			return reflect.ValueOf(bool2int(a.Bool())), b
		},
		{reflect.Bool, reflect.Int64}: func(a, b reflect.Value) (reflect.Value, reflect.Value) {
			return reflect.ValueOf(int64(bool2int(a.Bool()))), b
		},

		{reflect.Int, reflect.Float64}: func(a, b reflect.Value) (reflect.Value, reflect.Value) {
			return reflect.ValueOf(float64(a.Int())), b
		},
		{reflect.Int, reflect.Complex128}: func(a, b reflect.Value) (reflect.Value, reflect.Value) {
			return reflect.ValueOf(complex128(complex(float64(a.Int()), 0))), b
		},
		{reflect.Float64, reflect.Complex128}: func(a, b reflect.Value) (reflect.Value, reflect.Value) {
			return reflect.ValueOf(complex128(complex(a.Float(), 0))), b
		},
	}
}
