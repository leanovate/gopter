package gopter

import "testing"

type Properties struct {
	t          *testing.T
	parameters *CheckParameters
	props      map[string]Prop
	propNames  []string
}

func NewProperties(t *testing.T, parameters *CheckParameters) *Properties {
	if parameters == nil {
		parameters = DefaultCheckParameters()
	}
	return &Properties{
		t:          t,
		parameters: parameters,
		props:      make(map[string]Prop, 0),
		propNames:  make([]string, 0),
	}
}

func (p *Properties) Property(name string, prop Prop) {
	p.propNames = append(p.propNames, name)
	p.props[name] = prop
}
