package gopter

// Gen generator of arbitrary values
// Usually properties are checked by verifing a condition holds true for arbitrary input parameters
type Gen func(*GenParameters) *GenResult

// Sample generate a sample value
// Depending on the state of the RNG the generate might fail to provide a sample
func (g Gen) Sample() (interface{}, bool) {
	return g(DefaultGenParameters()).Retrieve()
}

func (g Gen) Map(f func(interface{}) interface{}) Gen {
	return func(genParams *GenParameters) *GenResult {
		result := g(genParams)
		value, ok := result.Retrieve()
		if ok {
			return &GenResult{
				Shrinker: NoShrinker,
				result:   f(value),
				Labels:   result.Labels,
			}
		}
		return &GenResult{
			Shrinker: NoShrinker,
			result:   nil,
			Labels:   result.Labels,
		}
	}
}
