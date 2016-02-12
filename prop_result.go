package gopter

type propStatus int

const (
	// PropProof THe property was proved (i.e. it is known to be correct and will be always true)
	PropProof propStatus = iota
	// PropTrue The property was true this time
	PropTrue
	// PropFalse The property was false this time
	PropFalse
	// PropUndecided The property has no clear outcome this time
	PropUndecided
	// PropError The property has generated an error
	PropError
)

// PropResult contains the result of a property
type PropResult struct {
	Status propStatus
	Error  error
	Args   []PropArg
	Labels []string
}

// NewPropResult create a PropResult with label
func NewPropResult(success bool, label string) *PropResult {
	if success {
		return &PropResult{
			Status: PropTrue,
			Labels: []string{label},
		}
	}
	return &PropResult{
		Status: PropFalse,
		Labels: []string{label},
	}
}

// Success checks if the result was successful
func (r *PropResult) Success() bool {
	return r.Status == PropTrue || r.Status == PropProof
}

// WithArgs adds argument descriptors to the PropResult for reporting
func (r *PropResult) WithArgs(args []PropArg) *PropResult {
	r.Args = args
	return r
}
