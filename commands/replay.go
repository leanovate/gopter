package commands

import (
	"github.com/leanovate/gopter"
)

// Replay a sequence of commands on a system for regression testing
func Replay(systemUnderTest SystemUnderTest, initialState State, commands ...Command) *gopter.PropResult {
	sequentialCommands := make([]shrinkableCommand, 0, len(commands))
	for _, command := range commands {
		sequentialCommands = append(sequentialCommands, shrinkableCommand{command: command, shrinker: gopter.NoShrinker})
	}
	actions := actions{
		initialStateProvider: func() State { return initialState },
		sequentialCommands:   sequentialCommands,
	}
	return actions.run(systemUnderTest)
}
