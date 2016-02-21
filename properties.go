package gopter

import (
	"fmt"
	"testing"
)

type Properties struct {
	parameters *TestParameters
	props      map[string]Prop
	propNames  []string
}

func NewProperties(parameters *TestParameters) *Properties {
	if parameters == nil {
		parameters = DefaultTestParameters()
	}
	return &Properties{
		parameters: parameters,
		props:      make(map[string]Prop, 0),
		propNames:  make([]string, 0),
	}
}

func (p *Properties) Property(name string, prop Prop) {
	p.propNames = append(p.propNames, name)
	p.props[name] = prop
}

func (p *Properties) Run(t *testing.T) {
	for _, propName := range p.propNames {
		prop := p.props[propName]

		result := prop.Check(p.parameters)

		if result.Passed() {
			fmt.Println(formatLines(fmt.Sprintf("+ %s: %s", propName, result.Report()), "", "", 75))
		} else {
			fmt.Println(formatLines(fmt.Sprintf("! %s: %s", propName, result.Report()), "", "", 75))
			t.Fail()
		}
	}
}
