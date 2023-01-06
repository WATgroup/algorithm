package algorithm

// Whether the supplied slice contains element 'e' or not
func Contains[T comparable](a []T, e T) bool {
	for _, x := range a {
		if e == x {
			return true
		}
	}
	return false
}
