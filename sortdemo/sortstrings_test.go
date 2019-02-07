package sortdemo

import (
	"reflect"
	"sort"
	"testing"
)

func createSlice() SSlice {
	return SSlice{"xyz", "abc", "aaa", "ab"}
}

func checkResult(t *testing.T, ss SSlice) {
	if !reflect.DeepEqual(ss, SSlice{"aaa", "ab", "abc", "xyz"}) {
		t.Errorf("Bad order %v", ss)
	}
}

func TestStringSort(t *testing.T) {
	ss := createSlice()
	sort.Sort(ss)
	checkResult(t, ss)
}

func TestStringSort2(t *testing.T) {
	ss := createSlice()
	sort.Sort(sort.StringSlice(ss))
	checkResult(t, ss)
}

func TestStringSort3(t *testing.T) {
	ss := createSlice()
	sort.Strings(ss)
	checkResult(t, ss)
}
