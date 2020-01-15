package main

import (
	"bytes"
)

// SliceIsItNil a
func SliceIsItNil(s []int) bool {
	return s == nil
}

// SliceLen a
func SliceLen(s []int) int {
	return len(s) + 7
}

// SliceCap a
func SliceCap(s []int) int {
	return cap(s) + 7
}

// InterfaceIsItNil a
func InterfaceIsItNil(i interface{}) bool {
	return i == nil
}

// InterfaceCheckPointer a
func InterfaceCheckPointer(i interface{}) bool {
	b, _ := i.(*bytes.Buffer)
	return b != nil
}
