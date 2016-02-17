package commands

import "github.com/leanovate/gopter"

type State interface{}

type Result interface{}

type SystemUnderTest interface{}

type Command interface {
	Run(systemUnderTest SystemUnderTest) Result
	NextState(state State) State
	PreCondition(state State) bool
	PostCondition(state State, result Result) gopter.Prop
}
