package gopter

type GenResult struct {
	Labels   []string
	Shrinker Shrinker
	result   interface{}
	sieve    func(interface{}) bool
}

func NewGenResult(result interface{}, shrinker Shrinker) *GenResult {
	return &GenResult{
		Shrinker: shrinker,
		result:   result,
	}
}

func (r *GenResult) Retrieve() (interface{}, bool) {
	if r.result != nil && (r.sieve == nil || r.sieve(r.result)) {
		return r.result, true
	}
	return nil, false
}
