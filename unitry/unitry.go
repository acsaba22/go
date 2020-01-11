package unitry

import (
	"fmt"
)

func playground() {
	s := "Grüezi"
	fmt.Println(s)
}

func lenUtf8(s string) int {
	// One unicode character can be multiple bytes.
	// Loop over the string with an integer and advance the int with the size of the current rune.
	// $ go doc utf8.DecodeRuneInString
	return 0
}

func lenUtf8Std(s string) int {
	// find lenght counter in utf8 package
	// go doc utf8
	return 0
}

func lenUtf8For(s string) int {
	// Use a range based for loop
	// https://golang.org/ref/spec#For_statements
	return 0
}

// works on utf8 too
func hasPrefix(s, prefix string) bool {
	return false
}

func contains(s, substr string) bool {
	return false
}

func intsToString(values []int) string {
	// go doc strconv.itoa
	return "[]"
}

func intsToStringFast(values []int) string {
	// go doc strings
	// go doc bytes
	// go doc bytes.Buffer
	// go doc fmt.Fprintf
	return "[]"
}
