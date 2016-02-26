package arbitrary

import (
	"reflect"
	"time"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

type Arbitrary struct {
	generators map[reflect.Type]gopter.Gen
}

func DefaultArbitrary() *Arbitrary {
	return &Arbitrary{
		generators: map[reflect.Type]gopter.Gen{
			reflect.TypeOf(time.Now()): gen.Time(),
		},
	}
}

func (a *Arbitrary) Gen(rt reflect.Type) gopter.Gen {
	if gen, ok := a.generators[rt]; ok {
		return gen
	}
	return a.genForKind(rt)
}
