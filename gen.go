package gopter

type Gen interface {
	DoApply(*GenParameters) GenResult
}
