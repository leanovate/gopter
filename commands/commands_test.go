package commands_test

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
)

type counter struct {
	value int
}

type counterCommands struct {
}

func (c *counterCommands) NewSystemUnderTest() commands.SystemUnderTest {
	return &counter{value: 0}
}

func (c *counterCommands) DestroySystemUnderTest(commands.SystemUnderTest) {
}

func (c *counterCommands) GenInitialState() gopter.Gen {
	return gen.Const(0)
}

func (c *counterCommands) GenCommand(state commands.State) gopter.Gen {
	return nil
}

func (c *counterCommands) InitialPreCondition(state commands.State) bool {
	return state.(int) == 0
}

func TestCommands(t *testing.T) {

}
