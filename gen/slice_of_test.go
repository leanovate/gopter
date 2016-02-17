package gen_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestSliceOf(t *testing.T) {
	genParams := gopter.DefaultGenParameters()
	genParams.Size = 50
	elementGen := gen.Const("element")
	sliceGen := gen.SliceOf(elementGen)

	for i := 0; i < 100; i++ {
		sample, ok := sliceGen(genParams).Retrieve()

		if !ok {
			t.Error("Sample was not ok")
		}
		strings, ok := sample.([]string)
		if !ok {
			t.Errorf("Sample not slice of string: %#v", sample)
		} else {
			if len(strings) >= 50 {
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

func TestSliceOfN(t *testing.T) {
	elementGen := gen.Const("element")
	sliceGen := gen.SliceOfN(10, elementGen)

	for i := 0; i < 100; i++ {
		sample, ok := sliceGen.Sample()

		if !ok {
			t.Error("Sample was not ok")
		}
		strings, ok := sample.([]string)
		if !ok {
			t.Errorf("Sample not slice of string: %#v", sample)
		} else {
			if len(strings) != 10 {
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
