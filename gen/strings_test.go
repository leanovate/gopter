package gen_test

import (
	"testing"
	"unicode"
	"unicode/utf8"

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
	commonGeneratorTest(t, "any string", gen.AnyString(), func(value interface{}) bool {
		str, ok := value.(string)

		if !ok {
			return false
		}
		for _, ch := range str {
			if !utf8.ValidRune(ch) {
				return false
			}
		}
		return true
	})
}

func TestAlphaString(t *testing.T) {
	commonGeneratorTest(t, "alpha string", gen.AlphaString(), func(value interface{}) bool {
		str, ok := value.(string)

		if !ok {
			return false
		}
		for _, ch := range str {
			if !utf8.ValidRune(ch) || !unicode.IsLetter(ch) {
				return false
			}
		}
		return true
	})
}

func TestNumString(t *testing.T) {
	commonGeneratorTest(t, "num string", gen.NumString(), func(value interface{}) bool {
		str, ok := value.(string)

		if !ok {
			return false
		}
		for _, ch := range str {
			if !utf8.ValidRune(ch) || !unicode.IsDigit(ch) {
				return false
			}
		}
		return true
	})
}

func TestIdentifier(t *testing.T) {
	commonGeneratorTest(t, "identifiers", gen.Identifier(), func(value interface{}) bool {
		str, ok := value.(string)

		if !ok {
			return false
		}
		if len(str) == 0 || !unicode.IsLetter([]rune(str)[0]) {
			return false
		}
		for _, ch := range str {
			if !utf8.ValidRune(ch) || (!unicode.IsDigit(ch) && !unicode.IsLetter(ch)) {
				return false
			}
		}
		return true
	})
}
