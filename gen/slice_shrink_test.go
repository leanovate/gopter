package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestSliceShrink(t *testing.T) {
	oneShrink := intSliceShinks(gen.SliceShrinker(gen.Int64Shrinker)([]int64{0}))
	if !intSliceSliceEquals(oneShrink, [][]int64{}) {
		t.Errorf("Invalid oneShrink: %#v", oneShrink)
	}

	twoShrink := intSliceShinks(gen.SliceShrinker(gen.Int64Shrinker)([]int64{0, 1}))
	if !intSliceSliceEquals(twoShrink, [][]int64{
		[]int64{1},
		[]int64{0},
		[]int64{0, 0},
	}) {
		t.Errorf("Invalid twoShrink: %#v", twoShrink)
	}

	threeShrink := intSliceShinks(gen.SliceShrinker(gen.Int64Shrinker)([]int64{0, 1, 2}))
	if !intSliceSliceEquals(threeShrink, [][]int64{
		[]int64{1, 2},
		[]int64{0, 2},
		[]int64{0, 1},
		[]int64{0, 0, 2},
		[]int64{0, 1, 0},
		[]int64{0, 1, 1},
		[]int64{0, 1, -1},
	}) {
		t.Errorf("Invalid threeShrink: %#v", threeShrink)
	}

	fourShrink := intSliceShinks(gen.SliceShrinker(gen.Int64Shrinker)([]int64{0, 1, 2, 3}))
	if !intSliceSliceEquals(fourShrink, [][]int64{
		[]int64{2, 3},
		[]int64{0, 1},
		[]int64{1, 2, 3},
		[]int64{0, 2, 3},
		[]int64{0, 1, 3},
		[]int64{0, 1, 2},
		[]int64{0, 0, 2, 3},
		[]int64{0, 1, 0, 3},
		[]int64{0, 1, 1, 3},
		[]int64{0, 1, -1, 3},
		[]int64{0, 1, 2, 0},
		[]int64{0, 1, 2, 2},
		[]int64{0, 1, 2, -2},
	}) {
		t.Errorf("Invalid fourShrink: %#v", fourShrink)
	}
}

func intSliceShinks(shrink gopter.Shrink) [][]int64 {
	result := make([][]int64, 0)

	value, ok := shrink()
	for ok {
		result = append(result, value.([]int64))
		value, ok = shrink()
	}
	return result
}

func intSliceSliceEquals(a, b [][]int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if !int64SliceEquals(e, b[i]) {
			return false
		}
	}
	return true
}
