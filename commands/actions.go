package commands

import (
	"fmt"
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

type actions struct {
	initialState       State
	sequentialCommands []Command
	// parallel commands will come later
}

func (a *actions) String() string {
	return fmt.Sprintf("initial=%v sequential=%s", a.initialState, a.sequentialCommands)
}

func (a *actions) run(systemUnderTest SystemUnderTest) (interface{}, error) {
	state := a.initialState
	propResult := &gopter.PropResult{Status: gopter.PropTrue}
	for _, command := range a.sequentialCommands {
		if !command.PreCondition(state) {
			return &gopter.PropResult{Status: gopter.PropFalse}, nil
		}
		result := command.Run(systemUnderTest)
		propResult = propResult.And(command.PostCondition(state, result))
		state = command.NextState(state)
	}
	return propResult, nil
}

type sizedCommands struct {
	state    State
	commands []Command
}

func actionsShrinker(v interface{}) gopter.Shrink {
	a := v.(*actions)
	return gen.SliceShrinker(gopter.NoShrinker)(a.sequentialCommands).Map(func(v []Command) *actions {
		return &actions{
			initialState:       a.initialState,
			sequentialCommands: v,
		}
	})
}

func genActions(commands Commands) gopter.Gen {
	return commands.GenInitialState().FlatMap(func(initialState interface{}) gopter.Gen {
		return genSizedCommands(commands, initialState.(State)).Map(func(v sizedCommands) *actions {
			return &actions{
				initialState:       initialState.(State),
				sequentialCommands: v.commands,
			}
		}).WithShrinker(actionsShrinker)
	}, reflect.TypeOf((*actions)(nil)))
}

func genSizedCommands(commands Commands, inistialState State) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		sizedCommandsGen := gen.Const(sizedCommands{
			state:    inistialState,
			commands: make([]Command, 0, genParams.Size),
		})
		for i := 0; i < genParams.Size; i++ {
			sizedCommandsGen = sizedCommandsGen.FlatMap(func(v interface{}) gopter.Gen {
				prev := v.(sizedCommands)
				return gen.RetryUntil(commands.GenCommand(prev.state), func(command Command) bool {
					return command.PreCondition(prev.state)
				}, 100).Map(func(command Command) sizedCommands {
					return sizedCommands{
						state:    command.NextState(prev.state),
						commands: append(prev.commands, command),
					}
				})
			}, reflect.TypeOf(sizedCommands{}))
		}
		return sizedCommandsGen(genParams)
	}
}
