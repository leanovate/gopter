package gen_test

import (
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestRune(t *testing.T) {
	commonGeneratorTest(t, "rune", gen.Rune(), func(value interface{}) bool {
		v, ok := value.(rune)
		return ok && utf8.ValidRune(v)
	})
}

func TestNumChar(t *testing.T) {
	commonGeneratorTest(t, "num char", gen.NumChar(), func(value interface{}) bool {
		v, ok := value.(rune)
		return ok && unicode.IsNumber(v)
	})
}

func TestAlphaUpper(t *testing.T) {
	commonGeneratorTest(t, "alpha upper char", gen.AlphaUpperChar(), func(value interface{}) bool {
		v, ok := value.(rune)
		return ok && unicode.IsUpper(v) && unicode.IsLetter(v)
	})
}

func TestAlphaLower(t *testing.T) {
	commonGeneratorTest(t, "alpha lower char", gen.AlphaLowerChar(), func(value interface{}) bool {
		v, ok := value.(rune)
		return ok && unicode.IsLower(v) && unicode.IsLetter(v)
	})
}

func TestAlphaChar(t *testing.T) {
	commonGeneratorTest(t, "alpha char", gen.AlphaChar(), func(value interface{}) bool {
		v, ok := value.(rune)
		return ok && unicode.IsLetter(v)
	})
}

func TestAnyString(t *testing.T) {
	alphaStrGen := gen.AnyString()
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
			if !utf8.ValidRune(ch) {
				t.Errorf("Invalid string: %#v", v)
			}
		}
	}
}

func TestAlphaString(t *testing.T) {
	alphaStrGen := gen.AlphaString()
	for i := 0; i < 100; i++ {
		result := alphaStrGen(gopter.DefaultGenParameters())
		value, ok := result.Retrieve()

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
		if result.Sieve == nil || result.Sieve("01") {
			t.Error("Invalid sieve")
		}
	}
}

func TestNumString(t *testing.T) {
	numStrGen := gen.NumString()
	for i := 0; i < 100; i++ {
		result := numStrGen(gopter.DefaultGenParameters())
		value, ok := result.Retrieve()

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
		if result.Sieve == nil || result.Sieve("abc") {
			t.Error("Invalid sieve")
		}
	}
}

func TestIdentifier(t *testing.T) {
	identifierGen := gen.Identifier()
	for i := 0; i < 100; i++ {
		result := identifierGen(gopter.DefaultGenParameters())
		value, ok := result.Retrieve()

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
		if result.Sieve == nil || result.Sieve("0ab") || result.Sieve("ab\n") {
			t.Error("Invalid sieve")
		}
	}
}
