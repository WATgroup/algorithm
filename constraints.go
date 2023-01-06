package algorithm

import (
	// 	"constraints"
	. "golang.org/x/exp/constraints"
)

type Number interface {
	Signed | Unsigned | Float //| Complex
}

type Enumerable interface {
	Integer | Float | ~string
}

type String interface {
	~string // | ~[]byte | ~[]rune
}
