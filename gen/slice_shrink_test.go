package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestSliceShrink(t *testing.T) {
	oneShrink := intSliceShinks(gen.SliceShrinker([]int{0}))
	if !intSliceSliceEquals(oneShrink, [][]int{}) {
		t.Errorf("Invalid oneShrink: %#v", oneShrink)
	}

}

func intSliceShinks(shrink gopter.Shrink) [][]int {
	result := make([][]int, 0)

	value, ok := shrink()
	for ok {
		result = append(result, value.([]int))
		value, ok = shrink()
	}
	return result
}

func intSliceSliceEquals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if !intSliceEquals(e, b[i]) {
			return false
		}
	}
	return true
}

func intSliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}
