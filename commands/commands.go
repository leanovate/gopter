package commands

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
)

type Commands interface {
	NewSystemUnderTest() SystemUnderTest
	DestroySystemUnderTest(SystemUnderTest)
	GenInitialState() gopter.Gen
	GenCommand(state State) gopter.Gen
	InitialPreCondition(state State) bool
}

func CommandsProp(commands Commands) gopter.Prop {
	return prop.ForAll1(genActions(commands), func(actions interface{}) (interface{}, error) {
		return true, nil
	})
}
