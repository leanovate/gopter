package gopter

type CheckStatus int

const (
	CheckPassed CheckStatus = iota
	CheckProved
	CheckFailed
	CheckExhausted
	CheckError
)

func (s CheckStatus) String() string {
	switch s {
	case CheckPassed:
		return "PASSED"
	case CheckProved:
		return "PROVED"
	case CheckFailed:
		return "FAILED"
	case CheckExhausted:
		return "EXHAUSTED"
	case CheckError:
		return "ERROR"
	}
	return ""
}

type CheckResult struct {
	Status    CheckStatus
	Succeeded int
	Discarded int
}

// Success checks if the result was successful
func (r *CheckResult) Success() bool {
	return r.Status == CheckPassed || r.Status == CheckProved
}
