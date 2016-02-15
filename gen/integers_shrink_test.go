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

	tenShrinks := int64Shinks(gen.Int64Shrinker(int64(10)))
	if !int64SliceEquals(tenShrinks, []int64{0, -5, 5, -8, 8, -9, 9}) {
		t.Errorf("Invalid tenShrinks: %#v", tenShrinks)
	}

	negTenShrinks := int64Shinks(gen.Int64Shrinker(int64(-10)))
	if !int64SliceEquals(negTenShrinks, []int64{0, 5, -5, 8, -8, 9, -9}) {
		t.Errorf("Invalid negTenShrinks: %#v", negTenShrinks)
	}

	leetShrink := int64Shinks(gen.Int64Shrinker(int64(1337)))
	if !int64SliceEquals(leetShrink, []int64{0, -669, 669, -1003, 1003, -1170, 1170, -1254, 1254, -1296, 1296, -1317, 1317, -1327, 1327, -1332, 1332, -1335, 1335, -1336, 1336}) {
		t.Errorf("Invalid leetShrink: %#v", leetShrink)
	}
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
