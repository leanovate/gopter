package arbitrary

import (
	"reflect"

	"github.com/leanovate/gopter"
)

type Arbitrary struct {
	generators map[reflect.Type]gopter.Gen
}

func DefaultArbitrary() *Arbitrary {
	return &Arbitrary{
		generators: map[reflect.Type]gopter.Gen{},
	}
}

func (a *Arbitrary) Gen(reflect.Type) gopter.Gen {
	return nil
}
