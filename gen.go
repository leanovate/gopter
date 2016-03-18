package gopter

import "reflect"

// Gen generator of arbitrary values
// Usually properties are checked by verifing a condition holds true for arbitrary input parameters
type Gen func(*GenParameters) *GenResult

// Sample generate a sample value
// Depending on the state of the RNG the generate might fail to provide a sample
func (g Gen) Sample() (interface{}, bool) {
	return g(DefaultGenParameters()).Retrieve()
}

// SuchThat creates a derived generator by adding a sieve, i.e. all generated values must have
// f(value) == true.
// Use this care, if the sieve to to fine the generator will have many misses which results
// in an undecided property.
func (g Gen) SuchThat(f func(interface{}) bool) Gen {
	return func(genParams *GenParameters) *GenResult {
		result := g(genParams)
		prevSieve := result.Sieve
		if prevSieve == nil {
			result.Sieve = f
		} else {
			result.Sieve = func(value interface{}) bool {
				return prevSieve(value) && f(value)
			}
		}
		return result
	}
}

// WithShrinker creates a derived generator with a specific shrinker
func (g Gen) WithShrinker(shrinker Shrinker) Gen {
	return func(genParams *GenParameters) *GenResult {
		result := g(genParams)
		if shrinker == nil {
			result.Shrinker = NoShrinker
		} else {
			result.Shrinker = shrinker
		}
		return result
	}
}

// Map creates a derived generators by mapping all generatored values with a given function.
// Note: The derived generator will not have a sieve or shrinker.
func (g Gen) Map(f func(interface{}) interface{}) Gen {
	return func(genParams *GenParameters) *GenResult {
		result := g(genParams)
		value, ok := result.Retrieve()
		if ok {
			mapped := f(value)
			return &GenResult{
				Shrinker:   NoShrinker,
				result:     mapped,
				Labels:     result.Labels,
				ResultType: reflect.TypeOf(mapped),
			}
		}
		return &GenResult{
			Shrinker:   NoShrinker,
			result:     nil,
			Labels:     result.Labels,
			ResultType: reflect.TypeOf(nil),
		}
	}
}

// FlatMap creates a derived generator by passing a generated value to a function which itself
// creates a generator.
func (g Gen) FlatMap(f func(interface{}) Gen) Gen {
	return func(genParams *GenParameters) *GenResult {
		result := g(genParams)
		value, ok := result.Retrieve()
		if ok {
			return f(value)(genParams)
		}
		return &GenResult{
			Shrinker:   NoShrinker,
			result:     nil,
			Labels:     result.Labels,
			ResultType: reflect.TypeOf(nil),
		}
	}
}

// CombineGens creates a generators from a list of generators.
// The result type will be a []interface{} containing the generated values of each generators in
// the list.
// Note: The combined generator will not have a sieve or shrinker.
func CombineGens(gens ...Gen) Gen {
	return func(genParams *GenParameters) *GenResult {
		values := make([]interface{}, len(gens))
		labels := make([]string, 0)
		shrinkers := make([]Shrinker, len(gens))
		sieves := make([]func(v interface{}) bool, len(gens))

		var ok bool
		for i, gen := range gens {
			result := gen(genParams)
			labels = append(labels, result.Labels...)
			shrinkers[i] = result.Shrinker
			sieves[i] = result.Sieve
			values[i], ok = result.Retrieve()
			if !ok {
				return &GenResult{
					Shrinker:   NoShrinker,
					result:     nil,
					Labels:     result.Labels,
					ResultType: reflect.TypeOf(values),
				}
			}
		}
		return &GenResult{
			Shrinker:   CombineShrinker(shrinkers...),
			result:     values,
			Labels:     labels,
			ResultType: reflect.TypeOf(values),
			Sieve: func(v interface{}) bool {
				values := v.([]interface{})
				for i, value := range values {
					if sieves[i] != nil && !sieves[i](value) {
						return false
					}
				}
				return true
			},
		}
	}
}
