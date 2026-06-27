// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package algorithm


// AllT is a pure-ternary-logic For-All operator, parameterized upon the
// "true", "false" and "neutral" values
// @Return the operator function
/** impl notes:
 * Note: Exhaustive operation (i.e. ∀x, r=r⊗x ) NOT needed... we can still
 * do shortcircuit evaluation as usual
 * N.W.: A naïve implementation would do a switch(x) for every value
 * ...but it turns out we don't need that much in this version :)
 */
func AllT[T comparable] (t,f,n T) func(input []T) T {

	return func(input []T) (ret T) {
		ret = n	// default is "dunno"
		for _,x := range input {
			if f == x {	// false
				return f	// as expected
			}
			if n == x {	// neutral
				continue	// assume no changes
			}
			if t == x {
				ret = t	// set return to "true"
			}
		}
		return // (ret)
	}
}


// AllTbool is a ternary-logic "For-All" operator
// most like the "pure" AllT, only limited to "true" or "false" return values
// It is indeed more performant than the "generalized" version
func AllTbool[V comparable] (t,f,n V) func(input []V) bool {
	// Implement the evaluation -- this is actually inlined by the compiler AFAIK
	return func(input[]V) bool {
		for _,x := range input {
			if f == x {
				return false
			}
			// just continue for 't' or 'n'
		}
		return true
	}
}


// AnyTbool is a ternary-logic "Any" operator
// most like the "pure" AnyT, only limited to "true" or "false" return values
// It is indeed more performant than the "generalized" version
func AnyTbool[V comparable] (t,f,n V) func(input []V) bool {
	// Implement the evaluation -- this is actually inlined by the compiler AFAIK
	return func(input[]V) bool {
		for _,x := range input {
			if t == x {
				return true
			}
			// just continue for 't' or 'n'
		}
		return false
	}
}

// AnyT is a pure-ternary-logic "Any" operator
// ∃x, x=t ⇔ ∀x, r=r⊕x  (r ~ r ∨ t)
// @Return the operator function
func AnyT[T comparable] (t,f,n T) func(input []T) T {

	return func(input[]T) (ret T) {
		ret = n
		for _,x := range input {
			if t == x {
				return t
			}
			if f == x {
				ret = f	// can be false
			}
			// if "dunno", just continue
		}
		return // ret
	}
}
