package sortdemo

import (
	"reflect"
	"sort"
	"testing"
)

func TestStringSort(t *testing.T) {
	ss := SSlice{"xyz", "abc", "aaa", "ab"}
	sort.Sort(ss)
	if !reflect.DeepEqual(ss, SSlice{"aaa", "ab", "abc", "xyz"}) {
		t.Errorf("Bad order %v", ss)
	}
}
