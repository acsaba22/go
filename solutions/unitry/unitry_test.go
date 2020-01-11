package unitry

import "testing"

func TestPlayground(t *testing.T) {
	playground()
}

func checkUtf8LenFunc(t *testing.T, f func(string) int) {
	e := []struct {
		s string
		l int
	}{
		{"", 0},
		{"abc", 3},
		{"a b c.", 6},
		{"Grüezi", 6},
		{"Hello мир!", 10},
	}
	for _, e := range e {
		if l := f(e.s); e.l != l {
			t.Errorf("len(%s) = %d (expected %d)", e.s, l, e.l)
		}
	}
}

func TestLenUtf8(t *testing.T) {
	checkUtf8LenFunc(t, lenUtf8)
}

func TestLenUtf8Std(t *testing.T) {
	checkUtf8LenFunc(t, lenUtf8Std)
}

func TestLenUtf8For(t *testing.T) {
	checkUtf8LenFunc(t, lenUtf8For)
}

func TestHasPrefix(t *testing.T) {
	e := []struct {
		s1 string
		s2 string
		b  bool
	}{
		{"", "", true},
		{"hi", "", true},
		{"hi", "h", true},
		{"hi", "i", false},
		{"hello world", "hello world", true},
		{"hello world", "hello w", true},
		{"hello world", "world", false},
		{"Grüezi", "Grüe", true},
		{"Grüezi", "Grue", false},
		{"short", "longer", false},
	}
	for _, e := range e {
		if b := hasPrefix(e.s1, e.s2); e.b != b {
			t.Errorf("hasPrefix(%s,%s) is %v", e.s1, e.s2, b)
		}
	}
}

func TestContains(t *testing.T) {
	e := []struct {
		s1 string
		s2 string
		b  bool
	}{
		{"", "", true},
		{"hi", "", true},
		{"hi", "h", true},
		{"hi", "i", true},
		{"hi", "x", false},
		{"hello world", "hello world", true},
		{"hello world", "hello w", true},
		{"hello world", "world", true},
		{"Grüezi", "rüe", true},
		{"Grüezi", "rue", false},
	}
	for _, e := range e {
		if b := contains(e.s1, e.s2); e.b != b {
			t.Errorf("contains(%s,%s) is %v", e.s1, e.s2, b)
		}
	}
}

func checkIntsToString(t *testing.T, f func([]int) string,
	fname string) {
	e := []struct {
		v []int
		s string
	}{
		{[]int{}, "[]"},
		{[]int{1}, "[1]"},
		{[]int{2, 1}, "[2, 1]"},
		{[]int{3, 2, 1}, "[3, 2, 1]"},
	}
	for _, e := range e {
		if s := f(e.v); e.s != s {
			t.Errorf("%s(%v) = %s (expected %s)",
				fname, e.v, s, e.s)
		}
	}

}

func TestIntsToString(t *testing.T) {
	checkIntsToString(t, intsToString, "intsToString")
}

func TestIntsToStringFast(t *testing.T) {
	checkIntsToString(t, intsToStringFast, "intsToStringFast")
}
