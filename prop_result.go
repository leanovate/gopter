package gopter

type status int

const (
	PropProof status = iota
	PropTrue
	PropFalse
	PropUndecided
	PropError
)

type PropResult struct {
	Status status
	Args   []PropArg
	Error  error
}
