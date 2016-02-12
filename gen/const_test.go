package gen_test

import (
	"testing"

	"github.com/untoldwind/gopter/gen"
)

func TestConstGen(t *testing.T) {
	constStr := gen.Const("some constant")

	for i := 0; i < 100; i++ {
		sample, ok := constStr.Sample()

		if !ok {
			t.Error("Sample was not ok")
		} else if sample.(string) != "some constant" {
			t.Errorf("Sample has wrong value: %#v", sample)
		}
	}
}
