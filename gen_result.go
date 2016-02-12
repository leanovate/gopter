package gopter

type GenResult struct {
	result interface{}
	Labels []string
	Sieve  func(interface{}) bool
}

func (r *GenResult) Retrieve() (interface{}, bool) {
	if r.result != nil && r.Sieve(r.result) {
		return r.result, true
	}
	return nil, false
}
