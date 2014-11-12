package spogoto

import (
	"math/rand"
)

// Interpreter interprets Spogoto code.
type Interpreter interface {
	RegisterStack(string, DataStack)
	Stack(string) DataStack
	Ok(string, int64) bool
	Bad(string, int64) bool
	RandInt() int64
}

type interpreter struct {
	DataStacks map[string]DataStack
	Rand       *rand.Rand
}

// RegisterStack adds a stack to the available DataStacks identified by name.
func (i *interpreter) RegisterStack(name string, stack DataStack) {
	i.DataStacks[name] = stack
}

// Stack returns the stack registered with name.
func (i *interpreter) Stack(name string) DataStack {
	s, ok := i.DataStacks[name]
	if !ok {
		s = &NullDataStack{}
	}

	return s
}

// Ok returns true if stack is available and has the number of elements
// required.
func (i *interpreter) Ok(name string, count int64) bool {
	return !i.Bad(name, count)
}

// Bad returns false if stack is available and has the number of elements
// required.
func (i *interpreter) Bad(name string, count int64) bool {
	return i.Stack(name).Lack(count)
}

func (i *interpreter) RandInt() int64 {
	return i.Rand.Int63n(10)
}

// NewInterpreter constructs a new Intepreter.
func NewInterpreter() *interpreter {
	return &interpreter{
		DataStacks: make(map[string]DataStack),
		Rand:       rand.New(rand.NewSource(rand.Int63())),
	}
}
