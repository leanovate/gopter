package gen_test

import (
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

func TestRune(t *testing.T) {
	runeGen := gen.Rune()
	for i := 0; i < 100; i++ {
		value, ok := runeGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid rune: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !utf8.ValidRune(v) {
			t.Errorf("Invalid rune: %#v", value)
		}
	}
}

func TestNumChar(t *testing.T) {
	numCharGen := gen.NumChar()
	for i := 0; i < 100; i++ {
		value, ok := numCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid numchar: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsNumber(v) {
			t.Errorf("Invalid numchar: %#v", value)
		}
	}
}

func TestAlphaUpper(t *testing.T) {
	alphaCharGen := gen.AlphaUpperChar()
	for i := 0; i < 100; i++ {
		value, ok := alphaCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid alphaupper: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsUpper(v) {
			t.Errorf("Invalid alphaupper: %#v", value)
		}
	}
}

func TestAlphaLower(t *testing.T) {
	alphaCharGen := gen.AlphaLowerChar()
	for i := 0; i < 100; i++ {
		value, ok := alphaCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid alphalower: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsLower(v) {
			t.Errorf("Invalid alphalower: %#v", value)
		}
	}
}

func TestAlphaChar(t *testing.T) {
	alphaCharGen := gen.AlphaChar()
	for i := 0; i < 100; i++ {
		value, ok := alphaCharGen.Sample()

		if !ok || value == nil {
			t.Errorf("Invalid alphachar: %#v", value)
		}
		v, ok := value.(rune)
		if !ok || !unicode.IsLetter(v) {
			t.Errorf("Invalid alphachar: %#v", value)
		}
	}
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
