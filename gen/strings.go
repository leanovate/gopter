package gen

import "github.com/untoldwind/gopter"

func RuneRange(min, max rune) gopter.Gen {
	return Int64Range(int64(min), int64(max)).Map(func(value interface{}) interface{} {
		return rune(value.(int64))
	}, nil).SuchThat(func(v interface{}) bool {
		return v.(rune) >= min && v.(rune) <= max
	})
}

func NumChar() gopter.Gen {
	return RuneRange('0', '9')
}

func AlphaUpperChar() gopter.Gen {
	return RuneRange('A', 'Z')
}

func AlphaLowerChar() gopter.Gen {
	return RuneRange('a', 'z')
}
