// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	f, err := os.Create("unary_ops_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprint(f, header)

	for _, fct := range []unaryDescr{
		{
			Name: "unaryInvert",
			GoOp: "^",
			PyOp: "~",
			Types: []typeDescr{
				// FIXME: boolDescr,
				intDescr, int8Descr, int16Descr, int32Descr, int64Descr,
				uintDescr, uint8Descr, uint16Descr, uint32Descr, uint64Descr,
			},
		},
		{
			Name: "unaryNegative",
			GoOp: "-",
			PyOp: "-",
			Types: []typeDescr{
				// FIXME: boolDescr,
				intDescr, int8Descr, int16Descr, int32Descr, int64Descr,
				uintDescr, uint8Descr, uint16Descr, uint32Descr, uint64Descr,
				float32Descr, float64Descr,
				complex64Descr, complex128Descr,
			},
		},
		{
			Name: "unaryNot",
			GoOp: "!",
			PyOp: "not",
			Types: []typeDescr{
				boolDescr,
				intDescr, int8Descr, int16Descr, int32Descr, int64Descr,
				uintDescr, uint8Descr, uint16Descr, uint32Descr, uint64Descr,
				float32Descr, float64Descr,
				complex64Descr, complex128Descr,
				stringDescr,
			},
		},
		{
			Name: "unaryPositive",
			GoOp: "+",
			PyOp: "+",
			Types: []typeDescr{
				// FIXME: boolDescr,
				intDescr, int8Descr, int16Descr, int32Descr, int64Descr,
				uintDescr, uint8Descr, uint16Descr, uint32Descr, uint64Descr,
				float32Descr, float64Descr,
				complex64Descr, complex128Descr,
			},
		},
	} {
		genUnary(f, fct)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

type typeDescr struct {
	Kind   reflect.Kind
	GoType reflect.Type
	GoMeth string
	Zero   string
}

type unaryDescr struct {
	Name  string
	GoOp  string
	PyOp  string
	Types []typeDescr
}

func genUnary(w io.Writer, desc unaryDescr) {
	fmt.Fprintf(
		w,
		"func %s(v Value) (Value, error) {\n\trv := reflect.ValueOf(v)\n\tswitch rv.Kind() {\n",
		desc.Name,
	)

	for _, typ := range desc.Types {
		switch desc.PyOp {
		case "not":
			fmt.Fprintf(
				w,
				"\tcase reflect.%v:\n\t\treturn %s(%s(rv.%s)==%s), nil\n",
				strings.Title(typ.Kind.String()),
				desc.GoOp,
				typ.GoType,
				typ.GoMeth,
				typ.Zero,
			)
		default:
			fmt.Fprintf(
				w,
				"\tcase reflect.%v:\n\t\treturn %s%s(rv.%s), nil\n",
				strings.Title(typ.Kind.String()),
				desc.GoOp,
				typ.GoType,
				typ.GoMeth,
			)
		}
	}

	fmt.Fprintf(
		w,
		"\t}\n\treturn nil, fmt.Errorf(\"TypeError: bad operand for unary %s: '%%v'\", rv.Kind())\n}\n\n",
		desc.PyOp,
	)
}

var (
	boolType       = reflect.TypeOf(false)
	byteType       = reflect.TypeOf(byte(0))
	intType        = reflect.TypeOf(int(0))
	int8Type       = reflect.TypeOf(int8(0))
	int16Type      = reflect.TypeOf(int16(0))
	int32Type      = reflect.TypeOf(int32(0))
	int64Type      = reflect.TypeOf(int64(0))
	uintType       = reflect.TypeOf(uint(0))
	uint8Type      = reflect.TypeOf(uint8(0))
	uint16Type     = reflect.TypeOf(uint16(0))
	uint32Type     = reflect.TypeOf(uint32(0))
	uint64Type     = reflect.TypeOf(uint64(0))
	float32Type    = reflect.TypeOf(float32(0))
	float64Type    = reflect.TypeOf(float64(0))
	complex64Type  = reflect.TypeOf(complex64(0))
	complex128Type = reflect.TypeOf(complex128(0))
	stringType     = reflect.TypeOf("")

	boolDescr       = typeDescr{reflect.Bool, boolType, "Bool()", "false"}
	byteDescr       = typeDescr{reflect.Uint8, byteType, "Uint()", "0"}
	uintDescr       = typeDescr{reflect.Uint, uintType, "Uint()", "0"}
	uint8Descr      = typeDescr{reflect.Uint8, uint8Type, "Uint()", "0"}
	uint16Descr     = typeDescr{reflect.Uint16, uint16Type, "Uint()", "0"}
	uint32Descr     = typeDescr{reflect.Uint32, uint32Type, "Uint()", "0"}
	uint64Descr     = typeDescr{reflect.Uint64, uint64Type, "Uint()", "0"}
	intDescr        = typeDescr{reflect.Int, intType, "Int()", "0"}
	int8Descr       = typeDescr{reflect.Int8, int8Type, "Int()", "0"}
	int16Descr      = typeDescr{reflect.Int16, int16Type, "Int()", "0"}
	int32Descr      = typeDescr{reflect.Int32, int32Type, "Int()", "0"}
	int64Descr      = typeDescr{reflect.Int64, int64Type, "Int()", "0"}
	float32Descr    = typeDescr{reflect.Float32, float32Type, "Float()", "0"}
	float64Descr    = typeDescr{reflect.Float64, float64Type, "Float()", "0"}
	complex64Descr  = typeDescr{reflect.Complex64, complex64Type, "Complex()", "0"}
	complex128Descr = typeDescr{reflect.Complex128, complex128Type, "Complex()", "0"}
	stringDescr     = typeDescr{reflect.String, stringType, "String()", `""`}
)

const header = `// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Automatically generated with 'go generate'.
// DO NOT MODIFY

package pygo

import (
	"fmt"
	"reflect"
)

`
