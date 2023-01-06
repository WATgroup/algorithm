package algorithm

// Diff returns the difference slice from two slices (values in s1 that are not in s2).
// Returns empty if given slice s1 is nil or empty.
// naïve implementation: O(n2) complexity
func Diff[T comparable](s1, s2 []T) []T {
	if len(s1) == 0 {
		return []T{}
	}

	var result []T
	for _, v := range s1 {
		if !Contains(s2, v) {
			result = append(result, v)
		}
	}
	return result
}

// Intersect returns the intersection of two slices.
// Returns empty if any of the given slices are nil or empty.
// naïve implementation: O(n2) complexity
func Intersect[T comparable](s1, s2 []T) []T {
	if 0 == len(s1) || 0 == len(s2) {
		return []T{}
	}

	var result []T
	for _, v := range s1 {
		if Contains(s2, v) {
			result = append(result, v)
		}
	}
	return result
}

////////////////////////////////////////////////////////////

// Select / Reject -> filter
// Supports all comparable types as slice values.
func Select[T, U comparable](slice []T, fn func(T) (bool,U)) []U {
	result := make([]U, 0, len(slice))
	for _, v := range slice {
		if ok,x := fn(v); ok {
			result = append(result, x)
		}
	}
	return result
}

func Reject[T, U comparable](slice []T, fn func(T) (bool,U)) []U {
	result := make([]U, 0, len(slice))
	for _, v := range slice {
		if ok,x := fn(v); !ok {
			result = append(result, x)
		}
	}
	return result
}


////////////////////////////////////////////////////////////

// Removes nil values from an array.
/* Example:
	var arr = []interface{}{1, 2, 3, nil, 4, 5}
	result := Compact(arr)  // [1, 2, 3, 4, 5]
*/
func Coalesce/*[T any]*/(input []any) []any {
	if nil == input { return input; }

	result := make([]any, 0, len(input))
	for _, v := range input {
		if nil != v {					// use reflection
			result = append(result, v)
		}
	}
	return result
}

////////////////////////////////////////////////////////////

func Unique[T comparable](slice []T) []T {

	if /*nil == slice ||*/ 1 > len(slice) {		// len(nil) -> 0
		return slice
	}

	// auxiliary map
	vals := make(map[T]bool,len(slice))

	// First pass: build unicity map
	for _, u := range slice {
		vals[u] = true
	}

	// Second pass: build result slice
	result := make([]T, 0, len(slice))
	for v,_ := range vals {
		result=append(result,v)
	}
	vals = nil		// attempt to free memory faster

	return result
}
