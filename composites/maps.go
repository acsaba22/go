package composites

import (
	"io"
)

// PrintSorted prints as many lines to the iowriter as the size of the map.
// Every output line should contains a key and a value separated by a space.
// The lines are sorted by keys.
func PrintSorted(m map[string]string, out io.Writer) {
	// Tips: create a key list, and use sort.Strings
	// fmt.Fprintf takes as first parameter an io.Writer
}

// Eq returns true if the two maps are equal.
func Eq(x, y map[string]int) bool {
	// Iterate over one check values in the other.
	return false
}
