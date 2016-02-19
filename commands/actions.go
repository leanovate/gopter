package commands

import (
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

type actions struct {
	state              State
	sequentialCommands []Command
	// parallel commands will come later
}

type sizedCommands struct {
	state    State
	commands []Command
}

func actionsShrinker(v interface{}) gopter.Shrink {
	return gen.SliceShrinker(gopter.NoShrinker)(v.(actions).sequentialCommands)
}

func genActions(commands Commands) gopter.Gen {
	return commands.GenInitialState().FlatMap(func(initialState interface{}) gopter.Gen {
		return genSizedCommands(commands, initialState.(State)).Map(func(v interface{}) interface{} {
			return actions{
				state:              initialState.(State),
				sequentialCommands: v.(sizedCommands).commands,
			}
		}).WithShrinker(actionsShrinker)
	})
}

func genSizedCommands(commands Commands, inistialState State) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		gen := gen.Const(sizedCommands{
			state:    inistialState,
			commands: make([]Command, 0, genParams.Size),
		})
		for i := 0; i < genParams.Size; i++ {
			gen = gen.FlatMap(func(v interface{}) gopter.Gen {
				prev := v.(sizedCommands)
				return commands.GenCommand(prev.state).SuchThat(func(command interface{}) bool {
					return command.(Command).PreCondition(prev.state)
				}).Map(func(command interface{}) interface{} {
					return sizedCommands{
						state:    command.(Command).NextState(prev.state),
						commands: append(prev.commands, command.(Command)),
					}
				})
			})
		}
		return gen(genParams)
	}
}
