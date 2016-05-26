package commands

import (
	"fmt"
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
)

type actions struct {
	// initialStateProvider has to reset/recreate the initial state exactly the
	// same every time.
	initialStateProvider func() State
	sequentialCommands   []Command
	// parallel commands will come later
}

func (a *actions) String() string {
	return fmt.Sprintf("initialState=%v sequential=%s", a.initialStateProvider(), a.sequentialCommands)
}

func (a *actions) run(systemUnderTest SystemUnderTest) (*gopter.PropResult, error) {
	state := a.initialStateProvider()
	propResult := &gopter.PropResult{Status: gopter.PropTrue}
	for _, command := range a.sequentialCommands {
		if !command.PreCondition(state) {
			return &gopter.PropResult{Status: gopter.PropFalse}, nil
		}
		result := command.Run(systemUnderTest)
		state = command.NextState(state)
		propResult = propResult.And(command.PostCondition(state, result))
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
			initialStateProvider: a.initialStateProvider,
			sequentialCommands:   v,
		}
	})
}

func genActions(commands Commands) gopter.Gen {
	genInitialState := commands.GenInitialState()
	genInitialStateProvider := gopter.Gen(func(params *gopter.GenParameters) *gopter.GenResult {
		seed := params.NextInt64()
		return gopter.NewGenResult(func() State {
			paramsWithSeed := params.CloneWithSeed(seed)
			if initialState, ok := genInitialState(paramsWithSeed).Retrieve(); ok {
				return initialState
			}
			return nil
		}, gopter.NoShrinker)
	}).SuchThat(func(initialStateProvoder func() State) bool {
		state := initialStateProvoder()
		return state != nil && commands.InitialPreCondition(state)
	})
	return genInitialStateProvider.FlatMap(func(v interface{}) gopter.Gen {
		initialStateProvider := v.(func() State)
		return genSizedCommands(commands, initialStateProvider).Map(func(v sizedCommands) *actions {
			return &actions{
				initialStateProvider: initialStateProvider,
				sequentialCommands:   v.commands,
			}
		}).SuchThat(func(actions *actions) bool {
			state := actions.initialStateProvider()
			for _, command := range actions.sequentialCommands {
				if !command.PreCondition(state) {
					return false
				}
				state = command.NextState(state)
			}
			return true
		}).WithShrinker(actionsShrinker)
	}, reflect.TypeOf((*actions)(nil)))
}

func genSizedCommands(commands Commands, initialStateProvider func() State) gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		sizedCommandsGen := gen.Const(sizedCommands{
			state:    initialStateProvider(),
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
