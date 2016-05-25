package commands_test

import (
	"fmt"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
)

// *****************************************
// Production code (i.e. the implementation)
// *****************************************

type Queue struct {
	inp  int
	outp int
	size int
	buf  []int
}

func New(n int) *Queue {
	return &Queue{
		inp:  0,
		outp: 0,
		size: n + 1,
		buf:  make([]int, n+1),
	}
}

func (q *Queue) Put(n int) int {
	q.buf[q.inp] = n
	q.inp = (q.inp + 1) % q.size
	return n
}

func (q *Queue) Get() int {
	ans := q.buf[q.outp]
	q.outp = (q.outp + 1) % q.size
	return ans
}

func (q *Queue) Size() int {
	return (q.inp - q.outp + q.size) % q.size
}

func (q *Queue) Init() {
	q.inp = 0
	q.outp = 0
}

// *****************************************
//               Test code
// *****************************************

// cbState holds the expected state (i.e. its the commands.State)
type cbState struct {
	size     int
	elements []int
}

func (c *cbState) String() string {
	return fmt.Sprintf("State(size=%d, elements=%v)", c.size, c.elements)
}

var getCommand = &commands.ProtoCommand{
	Name: "Get",
	RunFunc: func(q commands.SystemUnderTest) commands.Result {
		return q.(*Queue).Get()
	},
	NextStateFunc: func(state commands.State) commands.State {
		st := state.(*cbState)
		return &cbState{
			size:     st.size,
			elements: st.elements[1:],
		}
	},
	PreConditionFunc: func(state commands.State) bool {
		return len(state.(*cbState).elements) > 0
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		if result.(int) != 1 {
			return &gopter.PropResult{Status: gopter.PropFalse}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

var putCommand = &commands.ProtoCommand{
	Name: "Put",
	RunFunc: func(q commands.SystemUnderTest) commands.Result {
		return q.(*Queue).Put(1)
	},
	NextStateFunc: func(state commands.State) commands.State {
		st := state.(*cbState)
		return &cbState{
			size:     st.size,
			elements: append(st.elements, 1),
		}
	},
	PreConditionFunc: func(state commands.State) bool {
		s := state.(*cbState)
		return len(s.elements) < s.size
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		st := state.(*cbState)
		if result.(int) != st.elements[len(st.elements)-1] {
			return &gopter.PropResult{Status: gopter.PropFalse}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

var sizeCommand = &commands.ProtoCommand{
	Name: "Size",
	RunFunc: func(q commands.SystemUnderTest) commands.Result {
		return q.(*Queue).Size()
	},
	PreConditionFunc: func(state commands.State) bool {
		_, ok := state.(*cbState)
		return ok
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		if result.(int) != len(state.(*cbState).elements) {
			return &gopter.PropResult{Status: gopter.PropFalse}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

// cbCommands holds the expected state (i.e. its the commands.State)
type cbCommands struct {
	maxSize int
}

func NewCbCommands(maxSize int) *cbCommands {
	return &cbCommands{
		maxSize: maxSize,
	}
}

func (c *cbCommands) NewSystemUnderTest(initialState commands.State) commands.SystemUnderTest {
	s := initialState.(*cbState)
	q := New(s.size)
	for e := range s.elements {
		q.Put(e)
	}
	return q
}

func (c *cbCommands) DestroySystemUnderTest(sut commands.SystemUnderTest) {
	sut.(*Queue).Init()
}

func (c *cbCommands) GenInitialState() gopter.Gen {
	return gen.Const(&cbState{
		size:     c.maxSize,
		elements: make([]int, 0),
	})
}

func (c *cbCommands) InitialPreCondition(state commands.State) bool {
	s := state.(*cbState)
	return len(s.elements) >= 0 && len(s.elements) <= s.size
}

func (c *cbCommands) GenCommand(state commands.State) gopter.Gen {
	return gen.OneConstOf(getCommand, putCommand, sizeCommand)
}

// Kudos to @jamesd for providing this real world example
func Example_circularqueue() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("circular buffer", commands.Prop(NewCbCommands(10)))

	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// + circular buffer: OK, passed 100 tests.
}
