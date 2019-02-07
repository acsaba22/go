package composites

// RemoveEmpty removes empty strings from sslice.
// Modifies sslice and returns it.
func RemoveEmpty(sslice []string) []string {
	// return []string{}
	n := 0
	for i := 0; i < len(sslice); i++ {
		if sslice[i] != "" {
			sslice[n] = sslice[i]
			n++
		}
	}
	return sslice[:n]
}

// Remove the element with index v. Modifies v.
func Remove(v []int, i int) []int {
	// use built in function copy
	// return v
	copy(v[i:], v[i+1:])
	return v[:len(v)-1]
}
