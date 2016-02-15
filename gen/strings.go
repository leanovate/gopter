package gen

import "github.com/leanovate/gopter"

// RuneRange generates runes within a given range
func RuneRange(min, max rune) gopter.Gen {
	return Int64Range(int64(min), int64(max)).Map(func(value interface{}) interface{} {
		return rune(value.(int64))
	}, nil).SuchThat(func(v interface{}) bool {
		return v.(rune) >= min && v.(rune) <= max
	})
}

// NumChar generate arbitrary numberic character runes
func NumChar() gopter.Gen {
	return RuneRange('0', '9')
}

// AlphaUpperChar generate arbitrary uppercase alpha character runes
func AlphaUpperChar() gopter.Gen {
	return RuneRange('A', 'Z')
}

// AlphaLowerChar generate arbitrary lowercase alpha character runes
func AlphaLowerChar() gopter.Gen {
	return RuneRange('a', 'z')
}
