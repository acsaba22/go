package composites

import (
	"reflect"
	"strings"
	"testing"
)

func TestSliceRemoveEmpty(t *testing.T) {
	e := []struct {
		i []string
		o []string
	}{
		{[]string{}, []string{}},
		{[]string{"one", "", "three"}, []string{"one", "three"}},
		{[]string{"one", ""}, []string{"one"}},
		{[]string{"", "", "three"}, []string{"three"}},
		{[]string{"1", "", "2", "", "3"}, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3"}, []string{"1", "2", "3"}},
	}
	for _, e := range e {
		if o := RemoveEmpty(e.i); !reflect.DeepEqual(e.o, o) {
			t.Errorf("RemoveEmpty(%v) = [%v] (expected [%v])",
				strings.Join(e.i, ","),
				strings.Join(o, ","),
				strings.Join(e.o, ","))
		}
	}
}

func TestSliceRemoveEmptyModifiesOrig(t *testing.T) {
	s := []string{"one", "", "three"}
	RemoveEmpty(s)
	if s[1] != "three" {
		t.Errorf("RemoveEmpty expected to modify paramter.")
	}
}

func TestSliceRemove(t *testing.T) {
	v := []int{0, 1, 2, 3}

	v = Remove(v, 1)
	e := []int{0, 2, 3}
	if !reflect.DeepEqual(v, e) {
		t.Fatalf("RemoveEmpty([0 1 2 3], 1) = [%v] (expected %v)",
			v, e)
	}

	if v[1] != 2 {
		t.Fatalf("RemoveEmpty didn't modify the parameter")
	}

	v = Remove(v, 0)
	e = []int{2, 3}
	if !reflect.DeepEqual(v, e) {
		t.Fatalf("RemoveEmpty([0 2 3], 0) = [%v] (expected %v)",
			v, e)
	}

	v = Remove(v, 1)
	e = []int{2}
	if !reflect.DeepEqual(v, e) {
		t.Fatalf("RemoveEmpty([2 3], 1) = [%v] (expected %v)",
			v, e)
	}

	v = Remove(v, 0)
	e = []int{}
	if !reflect.DeepEqual(v, e) {
		t.Fatalf("RemoveEmpty([2], 0) = [%v] (expected %v)",
			v, e)
	}
}
