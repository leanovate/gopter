package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestInt64Shrink(t *testing.T) {
	zeroShrinks := int64Shinks(gen.Int64Shrinker(int64(0)))
	if !int64SliceEquals(zeroShrinks, []int64{}) {
		t.Errorf("Invalid zeroShrinks: %#v", zeroShrinks)
	}

	//	tenShrinks := int64Shinks(gen.Int64Shrinker(int64(10)))
	//	if !int64SliceEquals(tenShrinks, []int64{}) {
	//		t.Errorf("Invalid tenShrinks: %#v", tenShrinks)
	//	}
}

func int64Shinks(shrink gopter.Shrink) []int64 {
	result := make([]int64, 0)

	value, ok := shrink()
	for ok {
		result = append(result, value.(int64))
		value, ok = shrink()
	}
	return result
}

func int64SliceEquals(a, b []int64) bool {
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
