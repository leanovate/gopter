package gopter

type PrettyParameters struct {
	Verbosity int
}

func DefaultPrettyParameters() *PrettyParameters {
	return &PrettyParameters{
		Verbosity: 0,
	}
}

type Pretty interface {
	Apply(*PrettyParameters) string
}
