package composites

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestMapSortedPrint(t *testing.T) {
	m := map[string]string{
		"abc": "a",
		"xyz": "x",
		"iii": "i",
		"zzz": "z",
		"ccc": "ccc",
	}

	var out bytes.Buffer
	PrintSorted(m, &out)
	scanner := bufio.NewScanner(&out)
	scanner.Split(bufio.ScanWords)
	l := []string{}
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	actual := strings.Join(l, " ")
	expected := "abc a ccc ccc iii i xyz x zzz z"
	if actual != expected {
		t.Errorf("Print map in sorted order. Actual [%s] expected [%s]",
			actual, expected)
	}
}

func TestMapEq(t *testing.T) {
	type msi map[string]int
	e := []struct {
		x, y msi
		r    bool
	}{
		{msi{}, msi{}, true},
		{msi{"a": 1}, msi{"a": 1}, true},
		{msi{"a": 1}, msi{"a": 2}, false},
		{msi{"a": 1, "b": 0}, msi{"a": 1}, false},
		{msi{"a": 1}, msi{"a": 1, "b": 0}, false},
		{msi{"a": 1, "b": 0}, msi{"a": 1, "c": 0}, false},
	}
	for _, e := range e {
		if Eq(e.x, e.y) != e.r {
			t.Errorf("Eq(%v, %v) != %t", e.x, e.y, e.r)
		}
	}
}
