package composites

// RemoveEmpty removes empty strings from sslice.
// Modifies sslice and returns it.
func RemoveEmpty(sslice []string) []string {
	// return []string{}
	n := 0
	for _, v := range sslice {
		if v != "" {
			sslice[n] = v
			n++
		}
	}
	return sslice[:n]
}

// Remove the element with index k. Modifies v.
func Remove(v []int, k int) []int {
	// use built in function copy
	copy(v[k:], v[k+1:])
	return v[:len(v)-1]
}
