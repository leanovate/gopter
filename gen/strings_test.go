package gen_test

import (
	"testing"
	"unicode"

	"github.com/leanovate/gopter/gen"
)

func TestNumChar(t *testing.T) {
	numCharGen := gen.NumChar()
	for i := 0; i < 100; i++ {
		value, ok := numCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid char: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsNumber(v) {
			t.Errorf("Invalid char: %#v", value)
		}
	}
}

func TestAlphaUpper(t *testing.T) {
	alphaCharGen := gen.AlphaUpperChar()
	for i := 0; i < 100; i++ {
		value, ok := alphaCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid char: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsUpper(v) {
			t.Errorf("Invalid char: %#v", value)
		}
	}
}

func TestAlphaLower(t *testing.T) {
	alphaCharGen := gen.AlphaLowerChar()
	for i := 0; i < 100; i++ {
		value, ok := alphaCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid char: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsLower(v) {
			t.Errorf("Invalid char: %#v", value)
		}
	}
}
