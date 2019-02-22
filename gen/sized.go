package gen

import (
	"github.com/leanovate/gopter"
)

// Sized derives a generator from based on size
// This honors the `MinSize` and `MaxSize` of the `GenParameters` of the test suite.
func Sized(f func(int) gopter.Gen) gopter.Gen {
	return func(params *gopter.GenParameters) *gopter.GenResult {
		size := params.Rng.Intn(params.MaxSize-params.MinSize) + params.MinSize
		return f(size)(params)
	}
}
