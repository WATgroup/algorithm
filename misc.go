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

// Whether the element is within the supplied slice's values
// Actually, it's the same as above ... only with reversed params
func Constrain[T comparable](e T, a []T) bool {
	for _, x := range a {
		if e == x {
			return true
		}
	}
	return false
}
