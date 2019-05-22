package sortdemo

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestStringSort(t *testing.T) {
	s := []string{"xyz", "abc", "aaa", "ab"}
	sort.Sort(SSlice(s))
	if !reflect.DeepEqual(s, []string{"aaa", "ab", "abc", "xyz"}) {
		t.Errorf("Bad order %v", s)
	}
	fmt.Println(s)
}
