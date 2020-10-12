package prop

import (
	"reflect"
	"testing"
)

func TestDeepCloneScalars(t *testing.T) {
	i := 1234

	if cloned, ok := deepClone(reflect.ValueOf(i)).Interface().(int); !ok || cloned != i {
		t.Errorf("Invalid int clone: %#v != %#v", i, cloned)
	}

	f := float64(12.24)

	if cloned, ok := deepClone(reflect.ValueOf(f)).Interface().(float64); !ok || cloned != f {
		t.Errorf("Invalid float clone: %#v != %#v", f, cloned)
	}

	s := "Hallo"

	if cloned, ok := deepClone(reflect.ValueOf(s)).Interface().(string); !ok || cloned != s {
		t.Errorf("Invalid string clone: %#v != %#v", s, cloned)
	}
}

func TestDeepCloneArrays(t *testing.T) {
	arr1 := [2]string{"One", "Two"}

	if cloned, ok := deepClone(reflect.ValueOf(arr1)).Interface().([2]string); !ok || cloned[0] != "One" || cloned[1] != "Two" {
		t.Errorf("Invalid array clone: %#v != %#v", arr1, cloned)
	}

	arr2 := [3]int{111, 222, 333}

	if cloned, ok := deepClone(reflect.ValueOf(arr2)).Interface().([3]int); !ok || cloned[0] != 111 || cloned[1] != 222 || cloned[2] != 333 {
		t.Errorf("Invalid array clone: %#v != %#v", arr1, cloned)
	}
}
