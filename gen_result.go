package gopter

import "reflect"

type GenResult struct {
	Labels     []string
	Shrinker   Shrinker
	ResultType reflect.Type
	result     interface{}
	Sieve      func(interface{}) bool
}

func NewGenResult(result interface{}, shrinker Shrinker) *GenResult {
	return &GenResult{
		Shrinker:   shrinker,
		ResultType: reflect.TypeOf(result),
		result:     result,
	}
}

func NewEmptyResult(resultType reflect.Type) *GenResult {
	return &GenResult{
		ResultType: resultType,
		Shrinker:   NoShrinker,
	}
}

func (r *GenResult) Retrieve() (interface{}, bool) {
	if (r.Sieve == nil && r.result != nil) || (r.Sieve != nil && r.Sieve(r.result)) {
		return r.result, true
	}
	return nil, false
}
