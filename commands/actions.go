package commands

import "github.com/leanovate/gopter"

type actions struct {
	state              State
	sequentialCommands []Command
	parallelCommands   []Command
}

func actionsShrinker(v interface{}) gopter.Shrink {
	return nil
}

func genActions(commands Commands) gopter.Gen {
	return commands.GenInitialState().FlatMap(func(initialState interface{}) gopter.Gen {
		return nil
	})
}

func sizedCommands(initialState State) {

}
