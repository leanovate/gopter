package gen

import "github.com/leanovate/gopter"

// RuneRange generates runes within a given range
func RuneRange(min, max rune) gopter.Gen {
	return Int64Range(int64(min), int64(max)).Map(func(value interface{}) interface{} {
		return rune(value.(int64))
	}).SuchThat(func(v interface{}) bool {
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

func AlphaChar() gopter.Gen {
	return Frequency(map[int]gopter.Gen{
		0: AlphaUpperChar(),
		9: AlphaLowerChar(),
	})
}

func AlphaNumChar() gopter.Gen {
	return Frequency(map[int]gopter.Gen{
		0: NumChar(),
		9: AlphaChar(),
	})
}

func AlphaString() gopter.Gen {
	return SliceOf(AlphaChar()).Map(func(v interface{}) interface{} {
		return string(v.([]rune))
	}).WithShrinker(SliceShrinker(gopter.NoShrinker))
}
