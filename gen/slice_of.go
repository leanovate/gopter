package gen

import (
	"reflect"

	"github.com/leanovate/gopter"
)

// SliceOf generates an arbitrary slice of generated elements
func SliceOf(elementGen gopter.Gen) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		len := genParams.Rng.Intn(genParams.Size)
		element := elementGen(genParams)
		elementSieve := element.Sieve
		elementShrinker := element.Shrinker

		result := reflect.MakeSlice(element.ResultType, 0, len)

		for i := 0; i < len; i++ {
			value, ok := element.Retrieve()

			if ok {
				reflect.Append(result, reflect.ValueOf(value))
			}
			element = elementGen(genParams)
		}

		genResult := gopter.NewGenResult(result.Interface(), gopter.NoShrinker)
		genResult.Sieve = forAllSieve(elementSieve)
		genResult.Shrinker = SliceShrinker(elementShrinker)
		return genResult
	}
}

// SliceOfN generates a slice of generated elements with definied length
func SliceOfN(len int, elementGen gopter.Gen) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		element := elementGen(genParams)
		elementSieve := element.Sieve
		elementShrinker := element.Shrinker

		result := reflect.MakeSlice(element.ResultType, 0, len)
		for i := 0; i < len; i++ {
			value, ok := element.Retrieve()

			if ok {
				reflect.Append(result, reflect.ValueOf(value))
			}
			element = elementGen(genParams)
		}

		genResult := gopter.NewGenResult(result.Interface(), gopter.NoShrinker)
		genResult.Sieve = func(v interface{}) bool {
			rv := reflect.ValueOf(v)
			return rv.Len() == len && forAllSieve(elementSieve)(v)
		}
		genResult.Shrinker = SliceShrinkerOne(elementShrinker)
		return genResult
	}
}

func forAllSieve(elementSieve func(interface{}) bool) func(interface{}) bool {
	return func(v interface{}) bool {
		rv := reflect.ValueOf(v)
		for i := rv.Len() - 1; i >= 0; i-- {
			if !elementSieve(rv.Index(i)) {
				return false
			}
		}
		return true
	}
}
