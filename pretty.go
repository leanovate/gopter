package gopter

import "strings"

type PrettyParameters struct {
	Verbosity int
}

func DefaultPrettyParameters() *PrettyParameters {
	return &PrettyParameters{
		Verbosity: 0,
	}
}

func Pretty(prettyParams *PrettyParameters, result *TestResult) string {
	return ""
}

func formatLines(str, lead, trail string, width int) string {
	result := ""
	for _, line := range strings.Split(str, "\n") {
		result += breakLine(lead+line+trail, "  ", width)
	}
	return result
}

func breakLine(str, lead string, length int) string {
	result := ""
	for len(str) > length {
		result += lead + str[0:length] + "\n"
		str = str[length:]
	}
	result += str
	return result
}
