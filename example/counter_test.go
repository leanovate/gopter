package example

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
)

var GetCommand = &commands.ProtoCommand{
	Name: "GET",
	RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
		return systemUnderTest.(*Counter).Get()
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		if state.(int) != result.(int) {
			return &gopter.PropResult{Status: gopter.PropFalse}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

var IncCommand = &commands.ProtoCommand{
	Name: "INC",
	RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
		systemUnderTest.(*Counter).Inc()
		return nil
	},
	NextStateFunc: func(state commands.State) commands.State {
		return state.(int) + 1
	},
}

var DecCommand = &commands.ProtoCommand{
	Name: "DEC",
	RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
		systemUnderTest.(*Counter).Dec()
		return nil
	},
	NextStateFunc: func(state commands.State) commands.State {
		return state.(int) - 1
	},
}

var ResetCommand = &commands.ProtoCommand{
	Name: "RESET",
	RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
		systemUnderTest.(*Counter).Reset()
		return nil
	},
	NextStateFunc: func(state commands.State) commands.State {
		return 0
	},
}

type counterCommands struct {
}

func (c *counterCommands) NewSystemUnderTest() commands.SystemUnderTest {
	return &Counter{}
}

func (c *counterCommands) DestroySystemUnderTest(commands.SystemUnderTest) {
}

func (c *counterCommands) GenInitialState() gopter.Gen {
	return gen.Const(0)
}

func (c *counterCommands) InitialPreCondition(state commands.State) bool {
	return state.(int) == 0
}

func (c *counterCommands) GenCommand(state commands.State) gopter.Gen {
	return gen.OneConstOf(GetCommand, IncCommand, DecCommand)
}

func TestCounter(t *testing.T) {
	parameters := gopter.DefaultTestParameters()

	prop := commands.CommandsProp(&counterCommands{})

	result := prop.Check(parameters)

	if result.Passed() || len(result.Args) != 1 {
		t.Errorf("Expected to fail with args: %#v", result)
	}
	arg := result.Args[0]
	if arg.String() != "initial=0 sequential=[INC INC INC INC DEC GET]" {
		t.Errorf("Invalid arg in prop result: %v", arg)
	}
}
