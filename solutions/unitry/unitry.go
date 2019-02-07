package unitry

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

// TODO remove from tutorial, keep in solution
func playground() {
	s := "Grüezi"
	fmt.Println(s)
	for i, r := range s {
		fmt.Printf("%d %s ", i, string(r))
	}
	fmt.Println()
	fmt.Println(string(65))
}

func lenUtf8(s string) int {
	// go doc utf8.DecodeRuneInString
	// return 0
	i := 0
	n := 0
	for i < len(s) {
		_, k := utf8.DecodeRuneInString(s[i:])
		n++
		i += k
	}
	return n
}

func lenUtf8Std(s string) int {
	// find lenght counter in utf8 package
	// go doc utf8
	// return 0
	return utf8.RuneCountInString(s)
}

func lenUtf8For(s string) int {
	// Use a range based for loop
	// https://golang.org/ref/spec#For_statements
	n := 0
	for range s {
		n++
	}
	return n
}

// works on utf8 too
func hasPrefix(s, prefix string) bool {
	return len(prefix) <= len(s) && s[:len(prefix)] == prefix
}

func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func intsToString(values []int) string {
	// go doc strconv.itoa
	// return "[]"
	s := "["
	for i, v := range values {
		if 0 < i {
			s += ", "
		}
		s += strconv.Itoa(v)
	}
	s += "]"
	return s
}

func intsToStringFast(values []int) string {
	// go doc strings
	// go doc bytes
	// go doc bytes.Buffer
	// go doc fmt.Fprintf
	// return "[]"
	var b bytes.Buffer
	b.WriteByte('[')
	for i, v := range values {
		if 0 < i {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d", v)
	}
	b.WriteByte(']')
	return b.String()
}

// Extra possible assignments. Write test and functions:
// 1) basename: /usr/local/go/pkg/linux_amd64/fmt.a  => fmt
// 2) Format big numbers: 1234567 => 1,234,567
