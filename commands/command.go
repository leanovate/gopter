package commands

import "github.com/leanovate/gopter"

type State interface{}

type Result interface{}

type SystemUnderTest interface{}

type Command interface {
	Run(systemUnderTest SystemUnderTest) Result
	NextState(state State) State
	PreCondition(state State) bool
	PostCondition(state State, result Result) *gopter.PropResult
}

type ProtoCommand struct {
	RunFunc           func(systemUnderTest SystemUnderTest) Result
	NextStateFunc     func(state State) State
	PreConditionFunc  func(state State) bool
	PostConditionFunc func(state State, result Result) *gopter.PropResult
}

func (p *ProtoCommand) Run(systemUnderTest SystemUnderTest) Result {
	if p.RunFunc != nil {
		return p.RunFunc(systemUnderTest)
	}
	return nil
}

func (p *ProtoCommand) NextState(state State) State {
	if p.NextStateFunc != nil {
		return p.NextStateFunc(state)
	}
	return state
}

func (p *ProtoCommand) PreCondition(state State) bool {
	if p.PreConditionFunc != nil {
		return p.PreConditionFunc(state)
	}
	return true
}

func (p *ProtoCommand) PostCondition(state State, result Result) *gopter.PropResult {
	if p.PostConditionFunc != nil {
		return p.PostConditionFunc(state, result)
	}
	return &gopter.PropResult{Status: gopter.PropTrue}
}
