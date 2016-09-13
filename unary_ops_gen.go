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

func unaryInvert(v Value) (Value, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int:
		return ^int(rv.Int()), nil
	case reflect.Int8:
		return ^int8(rv.Int()), nil
	case reflect.Int16:
		return ^int16(rv.Int()), nil
	case reflect.Int32:
		return ^int32(rv.Int()), nil
	case reflect.Int64:
		return ^int64(rv.Int()), nil
	case reflect.Uint:
		return ^uint(rv.Uint()), nil
	case reflect.Uint8:
		return ^uint8(rv.Uint()), nil
	case reflect.Uint16:
		return ^uint16(rv.Uint()), nil
	case reflect.Uint32:
		return ^uint32(rv.Uint()), nil
	case reflect.Uint64:
		return ^uint64(rv.Uint()), nil
	}
	return nil, fmt.Errorf("TypeError: bad operand for unary ~: '%v'", rv.Kind())
}

func unaryNegative(v Value) (Value, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int:
		return -int(rv.Int()), nil
	case reflect.Int8:
		return -int8(rv.Int()), nil
	case reflect.Int16:
		return -int16(rv.Int()), nil
	case reflect.Int32:
		return -int32(rv.Int()), nil
	case reflect.Int64:
		return -int64(rv.Int()), nil
	case reflect.Uint:
		return -uint(rv.Uint()), nil
	case reflect.Uint8:
		return -uint8(rv.Uint()), nil
	case reflect.Uint16:
		return -uint16(rv.Uint()), nil
	case reflect.Uint32:
		return -uint32(rv.Uint()), nil
	case reflect.Uint64:
		return -uint64(rv.Uint()), nil
	case reflect.Float32:
		return -float32(rv.Float()), nil
	case reflect.Float64:
		return -float64(rv.Float()), nil
	case reflect.Complex64:
		return -complex64(rv.Complex()), nil
	case reflect.Complex128:
		return -complex128(rv.Complex()), nil
	}
	return nil, fmt.Errorf("TypeError: bad operand for unary -: '%v'", rv.Kind())
}

func unaryNot(v Value) (Value, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Bool:
		return !(bool(rv.Bool()) == false), nil
	case reflect.Int:
		return !(int(rv.Int()) == 0), nil
	case reflect.Int8:
		return !(int8(rv.Int()) == 0), nil
	case reflect.Int16:
		return !(int16(rv.Int()) == 0), nil
	case reflect.Int32:
		return !(int32(rv.Int()) == 0), nil
	case reflect.Int64:
		return !(int64(rv.Int()) == 0), nil
	case reflect.Uint:
		return !(uint(rv.Uint()) == 0), nil
	case reflect.Uint8:
		return !(uint8(rv.Uint()) == 0), nil
	case reflect.Uint16:
		return !(uint16(rv.Uint()) == 0), nil
	case reflect.Uint32:
		return !(uint32(rv.Uint()) == 0), nil
	case reflect.Uint64:
		return !(uint64(rv.Uint()) == 0), nil
	case reflect.Float32:
		return !(float32(rv.Float()) == 0), nil
	case reflect.Float64:
		return !(float64(rv.Float()) == 0), nil
	case reflect.Complex64:
		return !(complex64(rv.Complex()) == 0), nil
	case reflect.Complex128:
		return !(complex128(rv.Complex()) == 0), nil
	case reflect.String:
		return !(string(rv.String()) == ""), nil
	}
	return nil, fmt.Errorf("TypeError: bad operand for unary not: '%v'", rv.Kind())
}

func unaryPositive(v Value) (Value, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int:
		return +int(rv.Int()), nil
	case reflect.Int8:
		return +int8(rv.Int()), nil
	case reflect.Int16:
		return +int16(rv.Int()), nil
	case reflect.Int32:
		return +int32(rv.Int()), nil
	case reflect.Int64:
		return +int64(rv.Int()), nil
	case reflect.Uint:
		return +uint(rv.Uint()), nil
	case reflect.Uint8:
		return +uint8(rv.Uint()), nil
	case reflect.Uint16:
		return +uint16(rv.Uint()), nil
	case reflect.Uint32:
		return +uint32(rv.Uint()), nil
	case reflect.Uint64:
		return +uint64(rv.Uint()), nil
	case reflect.Float32:
		return +float32(rv.Float()), nil
	case reflect.Float64:
		return +float64(rv.Float()), nil
	case reflect.Complex64:
		return +complex64(rv.Complex()), nil
	case reflect.Complex128:
		return +complex128(rv.Complex()), nil
	}
	return nil, fmt.Errorf("TypeError: bad operand for unary +: '%v'", rv.Kind())
}
