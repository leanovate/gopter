package gopter

import "reflect"

// GenResult contains the result of a generator.
type GenResult struct {
	Labels     []string
	Shrinker   Shrinker
	ResultType reflect.Type
	result     interface{}
	Sieve      func(interface{}) bool
}

// NewGenResult creates a new generator result from for a concrete value and
// shrinker.
// Note: The concrete value "result" not be nil
func NewGenResult(result interface{}, shrinker Shrinker) *GenResult {
	return &GenResult{
		Shrinker:   shrinker,
		ResultType: reflect.TypeOf(result),
		result:     result,
	}
}

// NewEmptyResult creates an empty generator result.
// Unless the sieve does not explicitly allow it, empty (i.e. nil-valued)
// results are considered invalid.
func NewEmptyResult(resultType reflect.Type) *GenResult {
	return &GenResult{
		ResultType: resultType,
		Shrinker:   NoShrinker,
	}
}

// Retrieve gets the concrete generator result.
// If the result is invalid or does not pass the sieve there is no concrete
// value and the property using the generator should be undecided.
func (r *GenResult) Retrieve() (interface{}, bool) {
	if (r.Sieve == nil && r.result != nil) || (r.Sieve != nil && r.Sieve(r.result)) {
		return r.result, true
	}
	return nil, false
}

// RetrieveAsValue get the concrete generator result as reflect value.
// If the result is invalid or does not pass the sieve there is no concrete
// value and the property using the generator should be undecided.
func (r *GenResult) RetrieveAsValue() (reflect.Value, bool) {
	if r.result != nil && (r.Sieve == nil || r.Sieve(r.result)) {
		return reflect.ValueOf(r.result), true
	} else if r.result == nil && r.Sieve != nil && r.Sieve(r.result) {
		return reflect.Zero(r.ResultType), true
	}
	return reflect.Zero(r.ResultType), false
}
