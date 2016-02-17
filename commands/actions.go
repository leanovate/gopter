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

func genActions() gopter.Gen {
	return nil
}
