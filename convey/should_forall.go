package convey

import (
	"fmt"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
)

func ShouldSucceedForAll(condition interface{}, generators ...interface{}) string {
	parameters := gopter.DefaultTestParameters()
	gens := make([]gopter.Gen, len(generators))
	for i, generator := range generators {
		gen, ok := generator.(gopter.Gen)
		if !ok {
			return fmt.Sprintf("Expceted %#v to be a gropter.Gen", generator)
		}
		gens[i] = gen
	}

	property := prop.ForAll(condition, gens...)
	result := property.Check(parameters)

	if !result.Passed() {
		return fmt.Sprint(result)
	}
	return ""
}
