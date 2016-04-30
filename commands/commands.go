package commands

import (
	"reflect"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

// Commands provide an entry point for testing a stateful system
type Commands interface {
	// NewSystemUnderTest should create a new/isolated system under test
	NewSystemUnderTest() SystemUnderTest
	// DestroySystemUnderTest may perform any cleanup tasks to destroy a system
	DestroySystemUnderTest(SystemUnderTest)
	// GenInitialState provides a generator for the initial State
	GenInitialState() gopter.Gen
	// GenCommand provides a generator for applicable commands to for a state
	GenCommand(state State) gopter.Gen
	// InitialPreCondition checks if the initial state is valid
	InitialPreCondition(state State) bool
}

// ProtoCommands is a prototype implementation of the Commands interface
type ProtoCommands struct {
	NewSystemUnderTestFunc     func() SystemUnderTest
	DestroySystemUnderTestFunc func(SystemUnderTest)
	InitialStateGen            gopter.Gen
	GenCommandFunc             func(State) gopter.Gen
	InitialPreConditionFunc    func(State) bool
}

// NewSystemUnderTest should create a new/isolated system under test
func (p *ProtoCommands) NewSystemUnderTest() SystemUnderTest {
	if p.NewSystemUnderTestFunc != nil {
		return p.NewSystemUnderTestFunc()
	}
	return nil
}

// DestroySystemUnderTest may perform any cleanup tasks to destroy a system
func (p *ProtoCommands) DestroySystemUnderTest(systemUnderTest SystemUnderTest) {
	if p.DestroySystemUnderTestFunc != nil {
		p.DestroySystemUnderTestFunc(systemUnderTest)
	}
}

// GenCommand provides a generator for applicable commands to for a state
func (p *ProtoCommands) GenCommand(state State) gopter.Gen {
	if p.GenCommandFunc != nil {
		return p.GenCommandFunc(state)
	}
	return gen.Fail(reflect.TypeOf((*Command)(nil)).Elem())
}

// GenInitialState provides a generator for the initial State
func (p *ProtoCommands) GenInitialState() gopter.Gen {
	return p.InitialStateGen.SuchThat(func(state State) bool {
		return p.InitialPreCondition(state)
	})
}

// InitialPreCondition checks if the initial state is valid
func (p *ProtoCommands) InitialPreCondition(state State) bool {
	if p.InitialPreConditionFunc != nil {
		return p.InitialPreConditionFunc(state)
	}
	return true
}

// Prop creates a gopter.Prop from Commands
func Prop(commands Commands) gopter.Prop {
	return prop.ForAll1(genActions(commands), func(a interface{}) (interface{}, error) {
		systemUnderTest := commands.NewSystemUnderTest()
		defer commands.DestroySystemUnderTest(systemUnderTest)

		return a.(*actions).run(systemUnderTest)
	})
}
