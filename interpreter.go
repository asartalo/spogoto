package spogoto

// Interpreter interprets Spogoto code.
type Interpreter struct {
	DataStacks map[string]DataStack
}

// RegisterStack adds a stack to the available DataStacks identified by name.
func (i Interpreter) RegisterStack(name string, stack DataStack) {
	i.DataStacks[name] = stack
}

// Stack returns the stack registered with name.
func (i Interpreter) Stack(name string) DataStack {
	s, ok := i.DataStacks[name]
	if !ok {
		s = &NullDataStack{}
	}

	return s
}

// Ok returns true if stack is available and has the number of elements
// required.
func (i Interpreter) Ok(name string, count int64) bool {
	return !i.Bad(name, count)
}

// Bad returns false if stack is available and has the number of elements
// required.
func (i Interpreter) Bad(name string, count int64) bool {
	return i.Stack(name).Size() < count
}

// NewInterpreter constructs a new Intepreter.
func NewInterpreter() Interpreter {
	return Interpreter{make(map[string]DataStack)}
}
