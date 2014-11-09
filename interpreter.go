package spogoto

type Interpreter struct {
	DataStacks map[string]DataStack
}

func (i Interpreter) RegisterStack(name string, stack DataStack) {
	i.DataStacks[name] = stack
}

func (i Interpreter) Stack(name string) DataStack {
	s, ok := i.DataStacks[name]
	if !ok {
		s = &NullDataStack{}
	}

	return s
}

func (i Interpreter) Ok(name string, count int64) bool {
	return !i.Bad(name, count)
}

func (i Interpreter) Bad(name string, count int64) bool {
	return i.Stack(name).Size() < count
}

func NewInterpreter() Interpreter {
	return Interpreter{make(map[string]DataStack)}
}
