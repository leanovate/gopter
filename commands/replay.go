package commands

import (
	"github.com/leanovate/gopter"
)

// Replay a sequence of commands on a system for regression testing
func Replay(systemUnderTest SystemUnderTest, initialState State, commands ...Command) *gopter.PropResult {
	state := initialState
	propResult := &gopter.PropResult{Status: gopter.PropTrue}
	for _, command := range commands {
		if !command.PreCondition(state) {
			return &gopter.PropResult{Status: gopter.PropFalse}
		}
		result := command.Run(systemUnderTest)
		state = command.NextState(state)
		propResult = propResult.And(command.PostCondition(state, result))
	}
	return propResult
}
