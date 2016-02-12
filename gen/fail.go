package gen

import "github.com/untoldwind/gopter"

// Fail is a generator that always fails to generate a value
// Useful as fallback
var Fail = gopter.Gen(func(*gopter.GenParameters) *gopter.GenResult {
	return gopter.NewGenResult(nil, gopter.NoShrinker)
})
