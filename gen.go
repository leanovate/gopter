package gopter

// Gen generator of arbitrary values
// Usually properties are checked by verifing a condition holds true for arbitrary input parameters
type Gen func(*GenParameters) *GenResult

// Sample generate a sample value
// Depending on the state of the RNG the generate might fail to provide a sample
func (g Gen) Sample() (interface{}, bool) {
	return g(DefaultGenParameters).Retrieve()
}
