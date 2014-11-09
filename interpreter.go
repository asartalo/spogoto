package spogoto

type Interpreter struct {
	DataStacks map[string]DataStack
}

func (i Interpreter) Stack(name string) DataStack {
	s, ok := i.DataStacks[name]
	if !ok {
		s = &NullDataStack{}
	}

	return s
}
