package gen_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestArrayShrinkOne(t *testing.T) {
	oneShrink := gen.ArrayShrinkerOne(gen.Int64Shrinker)([1]int64{0}).All()
	if !reflect.DeepEqual(oneShrink, []interface{}{}) {
		t.Errorf("Invalid oneShrink: %#v", oneShrink)
	}

	threeShrink := gen.ArrayShrinkerOne(gen.Int64Shrinker)([3]int64{0, 1, 2}).All()
	if !reflect.DeepEqual(threeShrink, []interface{}{
		[3]int64{0, 0, 2},
		[3]int64{0, 1, 0},
		[3]int64{0, 1, 1},
		[3]int64{0, 1, -1},
	}) {
		t.Errorf("Invalid threeShrink: %#v", threeShrink)
	}

	fourShrink := gen.ArrayShrinkerOne(gen.Int64Shrinker)([4]int64{0, 1, 2, 3}).All()
	if !reflect.DeepEqual(fourShrink, []interface{}{
		[4]int64{0, 0, 2, 3},
		[4]int64{0, 1, 0, 3},
		[4]int64{0, 1, 1, 3},
		[4]int64{0, 1, -1, 3},
		[4]int64{0, 1, 2, 0},
		[4]int64{0, 1, 2, 2},
		[4]int64{0, 1, 2, -2},
	}) {
		t.Errorf("Invalid fourShrink: %#v", fourShrink)
	}
}
