package main

import "testing"

func TestFib(t *testing.T) {
	expectation := []struct {
		n int
		v int
	}{
		{0, 0},
		{1, 1},
		{3, 2},
		{5, 5},
		{10, 55},
	}
	for _, e := range expectation {
		if f := fib(e.n); e.v != f {
			t.Errorf("fib(%d) = %d (expected: %d)", e.n, f, e.v)
		}
	}
}
