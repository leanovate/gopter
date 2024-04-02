package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestArrayOfN(t *testing.T) {
	genParams := gopter.DefaultGenParameters()
	genParams.MaxSize = 50
	elementGen := gen.Const("element")
	arrayGen := gen.ArrayOfN(20, elementGen)

	for i := 0; i < 100; i++ {
		sample, ok := arrayGen(genParams).Retrieve()

		if !ok {
			t.Error("Sample was not ok")
		}
		strings, ok := sample.([20]string)
		if !ok {
			t.Errorf("Sample not slice of string: %#v", sample)
		} else {
			if len(strings) > 50 {
				t.Errorf("Sample has invalid length: %#v", len(strings))
			}
			for _, str := range strings {
				if str != "element" {
					t.Errorf("Sample contains invalid value: %#v", sample)
				}
			}
		}
	}
}
