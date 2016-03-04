package gen_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter/gen"
)

func TestOneConstOf(t *testing.T) {
	consts := gen.OneConstOf("one", "two", "three", "four")
	generated := make(map[string]bool, 0)
	for i := 0; i < 100; i++ {
		value, ok := consts.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid consts: %#v", value)
		}
		v, ok := value.(string)
		if !ok {
			t.Errorf("Invalid consts: %#v", value)
		}
		generated[v] = true
	}
	if !reflect.DeepEqual(generated, map[string]bool{
		"one":   true,
		"two":   true,
		"three": true,
		"four":  true,
	}) {
		t.Errorf("Not all consts where generated: %#v", generated)
	}
}

func TestOneGenOf(t *testing.T) {
	consts := gen.OneGenOf(gen.Const("one"), gen.Const("two"), gen.Const("three"), gen.Const("four"))
	generated := make(map[string]bool, 0)
	for i := 0; i < 100; i++ {
		value, ok := consts.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid consts: %#v", value)
		}
		v, ok := value.(string)
		if !ok {
			t.Errorf("Invalid consts: %#v", value)
		}
		generated[v] = true
	}
	if !reflect.DeepEqual(generated, map[string]bool{
		"one":   true,
		"two":   true,
		"three": true,
		"four":  true,
	}) {
		t.Errorf("Not all consts where generated: %#v", generated)
	}
}
