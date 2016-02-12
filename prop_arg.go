package gopter

type PropArg struct {
	Arg     interface{}
	OrigArg interface{}
	Label   string
	Shrinks int
}
