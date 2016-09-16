// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Automatically generated with 'go generate'.
// DO NOT MODIFY

package pygo

import "fmt"

func compareOp(a, b Value, op cmpOp) (Value, error) {
	switch op {
	case cmpOpLT:
	case cmpOpLE:
	case cmpOpEQ:
	case cmpOpNE:
	case cmpOpGT:
	case cmpOpGE:
	case cmpOpIN:
	case cmpOpNIN:
	case cmpOpIS:
	case cmpOpISNOT:
	case cmpOpISSUB:
	}
	// FIXME(sbinet)
	return nil, fmt.Errorf("pygo: internal error: no such compare-op %d", int(op))
}
