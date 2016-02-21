package gopter

import (
	"fmt"
	"strings"
)

// PropArg contains information about the specific values for a certain property check
// This is mostly used for reporting when a property has falsified
type PropArg struct {
	Arg     interface{}
	OrigArg interface{}
	Label   string
	Shrinks int
}

func (p *PropArg) Report(idx int) string {
	label := p.Label
	if label == "" {
		label = fmt.Sprintf("ARG_%d", idx)
	}
	result := fmt.Sprintf("%s: %v", label, p.Arg)
	if p.Shrinks > 0 {
		result += fmt.Sprintf("\n%s_ORIGINAL (%d shrinks): %v", label, p.Shrinks, p.OrigArg)
	}

	return result
}

type PropArgs []*PropArg

func (p PropArgs) Report() string {
	result := ""
	for i, arg := range p {
		if result != "" {
			result += "\n"
		}
		result += arg.Report(i)
	}
	return result
}

func NewPropArg(genResult *GenResult, shrinks int, value, origValue interface{}) *PropArg {
	return &PropArg{
		Label:   strings.Join(genResult.Labels, ", "),
		Arg:     value,
		OrigArg: origValue,
		Shrinks: shrinks,
	}
}
