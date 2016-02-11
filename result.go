package gopter

type status int

const (
	Proof status = iota
	True
	False
	Undecided
	Error
)

type Result struct {
	Status status
	Error  error
}
