package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestPtrOf(t *testing.T) {
	genParams := gopter.DefaultGenParameters()
	elementGen := gen.Const("element")
	ptrGen := gen.PtrOf(elementGen)

	for i := 0; i < 100; i++ {
		sample, ok := ptrGen(genParams).Retrieve()

		if !ok {
			t.Error("Sample was not ok")
		}
		if sample == nil {
			continue
		}
		stringPtr, ok := sample.(*string)
		if !ok {
			t.Errorf("Sample not pointer to string: %#v", sample)
		} else if *stringPtr != "element" {
			t.Errorf("Sample contains invalid value: %#v %#v", sample, *stringPtr)
		}
	}
}
