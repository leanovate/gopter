package gopter

type CheckStatus int

const (
	CheckPassed CheckStatus = iota
	CheckProved
	CheckFailed
	CheckExhausted
	CheckError
)

type CheckResult struct {
	Status    CheckStatus
	Succeeded int
	Discarded int
}
