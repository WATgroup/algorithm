package algorithm

// MergeDict: merges two dictionaries;
// destination values will get overwritten by those on src
func MergeDict[T any](dst map[string]T, src map[string]T) {

	if nil == src || nil == dst {
		return
	}
	if &dst == &src {
		return
	}

	for k, v := range src {
		dst[k] = v
	}
}
