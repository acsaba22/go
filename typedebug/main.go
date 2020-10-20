package main

import (
	"os"
)

func main() {
	s := []int{}
	r := 0
	if SliceIsItNil(s) {
		r = SliceLen(s) + SliceCap(s)
	} else if InterfaceIsItNil(s) {
		r = 1
	}
	os.Exit(r)
}
