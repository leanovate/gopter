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

func TestAlphaChar(t *testing.T) {
	alphaCharGen := gen.AlphaChar()
	for i := 0; i < 100; i++ {
		value, ok := alphaCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid char: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsLetter(v) {
			t.Errorf("Invalid char: %#v", value)
		}
	}
}

func TestAlphaString(t *testing.T) {
	alphaStrGen := gen.AlphaString()
	for i := 0; i < 100; i++ {
		value, ok := alphaStrGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid string: %#v", value)
		}
		v, ok := value.(string)
		if !ok {
			t.Errorf("Invalid string: %#v", value)
		}
		for _, ch := range v {
			if !unicode.IsLetter(ch) {
				t.Errorf("Invalid string: %#v", v)
			}
		}
	}
}

func TestNumString(t *testing.T) {
	numStrGen := gen.NumString()
	for i := 0; i < 100; i++ {
		value, ok := numStrGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid string: %#v", value)
		}
		v, ok := value.(string)
		if !ok {
			t.Errorf("Invalid string: %#v", value)
		}
		for _, ch := range v {
			if !unicode.IsDigit(ch) {
				t.Errorf("Invalid string: %#v", v)
			}
		}
	}
}

func TestIdentifier(t *testing.T) {
	identifierGen := gen.Identifier()
	for i := 0; i < 100; i++ {
		value, ok := identifierGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid string: %#v", value)
		}
		v, ok := value.(string)
		if !ok {
			t.Errorf("Invalid string: %#v", value)
		}
		if len(v) == 0 {
			t.Errorf("Invalid string: %#v", v)
		}
		if !unicode.IsLower([]rune(v)[0]) {
			t.Errorf("Invalid string: %#v", v)
		}
		for _, ch := range v {
			if !unicode.IsDigit(ch) && !unicode.IsLetter(ch) {
				t.Errorf("Invalid string: %#v", v)
			}
		}
	}
}
