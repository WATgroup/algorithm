// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>


package algorithm_test

// separate package to avoid polluting the package namespace with our "ternary test" type


import (
	. "github.com/WATgroup/algorithm"
	"testing"
)


type ternary uint32		// suggested ternary type
type ter = ternary
const (
	vTRUE ternary = (1<<32-1)	// -1 grrr
	vFALSE ternary = 0
	vDUNNO ternary = 0x80000000	// "-0"
)

type terTest struct {
	input []ternary
	name	string
	expected ternary
}


func TestAllT(t *testing.T) {

	// Setup....
	fAll := AllT[ternary](vTRUE,vFALSE,vDUNNO)

	tests := []terTest{{nil,"AllT·nil", vDUNNO},
		{[]ter{vTRUE,vDUNNO,vTRUE}, "AllT·dunno", vTRUE},
		{[]ter{vTRUE,vDUNNO,vFALSE}, "AllT·false", vFALSE}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != fAll(tc.input) {
				t.Fail()
			}
		})
	}
}

func TestAnyT(t *testing.T) {

	// Setup....
	fAny := AnyT[ternary](vTRUE,vFALSE,vDUNNO)

	tests := []terTest{{nil,"AnyT·nil", vDUNNO},
	{[]ter{vTRUE,vDUNNO,vTRUE}, "AnyT·dunno", vTRUE},
	{[]ter{vFALSE,vDUNNO,vTRUE}, "AnyT·dunno2", vTRUE},
	{[]ter{vTRUE,vDUNNO,vFALSE}, "AnyT·false", vTRUE}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != fAny(tc.input) {
				t.Fail()
			}
		})
	}
}
