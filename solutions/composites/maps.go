package composites

import (
	"fmt"
	"io"
	"sort"
)

// PrintSorted prints as many lines to the iowriter as the size of the map.
// Every output line should contains a key and a value separated by a space.
// The lines are sorted by keys.
func PrintSorted(m map[string]string, out io.Writer) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(out, "%s %s\n", k, m[k])
	}
}

// Eq returns true if the two maps are equal.
func Eq(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if v2, ok := y[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}
