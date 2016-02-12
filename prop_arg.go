package gopter

// PropArg contains information about the specific values for a certain property check
// This is mostly used for reporting when a property has falsified
type PropArg struct {
	Arg     interface{}
	OrigArg interface{}
	Label   string
	Shrinks int
}
