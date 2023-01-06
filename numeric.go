package algorithm


// numeric algorithms
// Supports all types which satisfy the "Number" type constraint:
// int/int8/int16/int32/int64, uint/uint8/uint16/uint32/uint64 float32, float64

// Sum sums the values of slice.
func Sum[T Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Min get the min value of slice.
// Returns 0 if slice is nil or empty.
func Min[T Number](slice []T) T {
	if 0 == len(slice) {
		return T(0)
	}

	var min T = slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min
}

// Max get the max value of slice.
// Returns 0 if slice is nil or empty.
func Max[T Number](slice []T) T {
	if 0 == len(slice) {
		return T(0)
	}

	var max T = slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
