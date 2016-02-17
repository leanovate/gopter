package commands

import "github.com/leanovate/gopter"

type Commands interface {
	NewSystemUnderTest() SystemUnderTest
	DestroySystemUnderTest() SystemUnderTest
	GenInitialState() gopter.Gen
	GenCommand() gopter.Gen
	InitialPreCondition(state State) bool
}

func CommandsProp(commands Commands) gopter.Prop {
	return nil
}
