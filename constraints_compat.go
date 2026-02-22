//go:build go1.18 || go1.19

package algorithm

import (
	//revive:disable-next-line:dot-imports
	. "golang.org/x/exp/constraints"
)

// type Complex interface {
// 	~complex64 | ~complex128
// }
//
// type Float interface {
// 	~float32 | ~float64
// }
//
// type Signed interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64
// }
//
// type Unsigned interface {
// 	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
// }
//
// type Integer interface {
// 	Signed | Unsigned
// }
//
// type Ordered interface {
// 	Integer | Float | ~string
// }

// The Int constraint (similar to Number) includes signed and unsigned ints
type Int interface {
	Signed | Unsigned
}
