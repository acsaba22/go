package main

import (
	"os"

	"mycompany.com/firstgo/typedebug"
)

func main() {
	s := []int{}
	r := 0
	if SliceIsItNil(s) {
		r = SliceLen(s) + typedebug.SliceCap(s)
	} else if InterfaceIsItNil(s) {
		r = 1
	}
	os.Exit(r)
}
