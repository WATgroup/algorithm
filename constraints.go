// Package algorithm provides some generic, functional-like primitives and
// extends "typeset" constraints for use in other libraries
package algorithm

import (
	// 	"constraints"
	//revive:disable-next-line:dot-imports
	. "golang.org/x/exp/constraints"
)

// The Number constraint: signed, unsigned or floating point
type Number interface {
	Signed | Unsigned | Float //| Complex
}

// Enumerable constraint: number or string-like
type Enumerable interface {
	Integer | Float | ~string
}

// A String -alike constraint
// Actually, byte slices are better optimized by the compiler...
type String interface {
	~string // | ~[]byte | ~[]rune
}
