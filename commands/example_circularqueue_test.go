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
	if q.inp == 0 { // Intentional bug to find
		q.buf[q.size-1] *= (n + 1)
	}
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
	size         int
	elements     []int
	takenElement int
}

func (st *cbState) TakeFront() {
	st.takenElement = st.elements[0]
	st.elements = append(st.elements[:0], st.elements[1:]...)
}

func (st *cbState) PushBack(value int) {
	st.elements = append(st.elements, value)
}

func (c *cbState) String() string {
	return fmt.Sprintf("State(size=%d, elements=%v)", c.size, c.elements)
}

var genGetCommand = gen.Const(&commands.ProtoCommand{
	Name: "Get",
	RunFunc: func(q commands.SystemUnderTest) commands.Result {
		return q.(*Queue).Get()
	},
	NextStateFunc: func(state commands.State) commands.State {
		state.(*cbState).TakeFront()
		return state
	},
	PreConditionFunc: func(state commands.State) bool {
		return len(state.(*cbState).elements) > 0
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		if result.(int) != state.(*cbState).takenElement {
			return &gopter.PropResult{Status: gopter.PropFalse}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
})

type putCommand int

func (value putCommand) Run(q commands.SystemUnderTest) commands.Result {
	return q.(*Queue).Put(int(value))
}
func (value putCommand) NextState(state commands.State) commands.State {
	state.(*cbState).PushBack(int(value))
	return state
}

func (putCommand) PreCondition(state commands.State) bool {
	s := state.(*cbState)
	return len(s.elements) < s.size
}

func (putCommand) PostCondition(state commands.State, result commands.Result) *gopter.PropResult {
	st := state.(*cbState)
	if result.(int) != st.elements[len(st.elements)-1] {
		return &gopter.PropResult{Status: gopter.PropFalse}
	}
	return &gopter.PropResult{Status: gopter.PropTrue}
}

func (value putCommand) String() string {
	return fmt.Sprintf("Put(%d)", value)
}

var genPutCommand = gen.Int().Map(func(value int) commands.Command {
	return putCommand(value)
}).WithShrinker(func(v interface{}) gopter.Shrink {
	return gen.IntShrinker(int(v.(putCommand))).Map(func(value int) putCommand {
		return putCommand(value)
	})
})

var genSizeCommand = gen.Const(&commands.ProtoCommand{
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
})

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
	return gen.IntRange(1, c.maxSize).Map(func(maxSize int) *cbState {
		return &cbState{
			size:     maxSize,
			elements: make([]int, 0, maxSize),
		}
	})
}

func (c *cbCommands) InitialPreCondition(state commands.State) bool {
	s := state.(*cbState)
	return len(s.elements) >= 0 && len(s.elements) <= s.size
}

func (c *cbCommands) GenCommand(state commands.State) gopter.Gen {
	return gen.OneGenOf(genGetCommand, genPutCommand, genSizeCommand)
}

// Kudos to @jamesd for providing this real world example.
// ... of course he did not implemented the bug, that was evil me
func Example_circularqueue() {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("circular buffer", commands.Prop(NewCbCommands(100)))

	// When using testing.T you might just use: properties.TestingRun(t)
	properties.Run(gopter.ConsoleReporter(false))
	// Output:
	// ! circular buffer: Falsified after 33 passed tests.
	// ARG_0: initialState=State(size=7, elements=[]) sequential=[Put(0) Put(0)
	//    Put(0) Put(0) Put(0) Get Put(0) Get Get Get Put(0) Get Put(1) Get Get Get]
	// ARG_0_ORIGINAL (48 shrinks): initialState=State(size=7, elements=[])
	//    sequential=[Put(-106526931) Get Size Size Put(-1590798911) Size
	//    Put(1121470879) Size Put(2086210077) Size Put(920967946) Put(-1336336465)
	//    Get Put(-1420016806) Get Get Get Put(1371806167) Get Size Put(556302804)
	//    Size Put(1154954099) Size Get Size Size Get Get Size Get Put(126492399)
	//    Size]
}
