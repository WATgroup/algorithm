package algorithm

// "Any" returns true if the given function returns true for any element of the slice.
// Supports all comparable types as slice values.
func Any[T comparable](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

// "All" returns true if the given function returns true for all elements of the slice.
// Supports all comparable types for slice values.
func All[T comparable](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter filters the values of slice by the given function.
// Returns empty if given slice is nil or empty.
func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

////////////////////////////////////////////////////////////////////////////////

// fn can't be nil
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, 0, len(slice))
	for _, v := range slice {
		result = append(result, fn(v))
	}
	return result
}

////////////////////

func Reduce[T comparable, U any] (slice []T, initial U, reducer func(U,T) U) U {
	if nil == slice || nil == reducer {
		return initial
	}

	var accum U = initial
	for _, v := range slice {
		accum = reducer(accum,v)
	}
	return accum
}
