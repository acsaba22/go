package composites

import (
	"io"
)

// PrintSorted prints as many lines to the iowriter as the size of the map.
// Every output line should contains a key and a value separated by a space.
// The lines are sorted by keys.
func PrintSorted(m map[string]string, out io.Writer) {
}

// Eq returns true if the two maps are equal.
func Eq(x, y map[string]int) bool {
	return false
}
