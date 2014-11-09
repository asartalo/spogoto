package spogoto

type FunctionMap map[string]func(DataStack, Interpreter)

// DataStack is a Stack used by the Interpreter
// to store data and has functions that can manipulate
// data from the interpreter
type DataStack interface {
	Stack
	Functions() FunctionMap
	Call(string, Interpreter)
}

type datastack struct {
	stack
	FunctionMap FunctionMap
}

func NewDataStack(elements Elements, functions FunctionMap) *datastack {
	d := &datastack{stack{elements}, functions}
	addCommonFunctions(d)
	return d
}

func addCommonFunctions(ds *datastack) {

	ds.FunctionMap["pop"] = func(d DataStack, i Interpreter) {
		d.Pop()
	}

	ds.FunctionMap["swap"] = func(d DataStack, i Interpreter) {
		d.Swap()
	}

	ds.FunctionMap["rotate"] = func(d DataStack, i Interpreter) {
		d.Rotate()
	}

	ds.FunctionMap["shove"] = func(d DataStack, i Interpreter) {
		if i.Bad("integer", 1) {
			return
		}

		idx := i.Stack("integer").Pop().(int64)
		d.Shove(d.Pop(), idx)
	}

	ds.FunctionMap["yank"] = func(d DataStack, i Interpreter) {
		if i.Bad("integer", 1) {
			return
		}

		idx := i.Stack("integer").Pop().(int64)
		d.Yank(idx)
	}

	ds.FunctionMap["yankdup"] = func(d DataStack, i Interpreter) {
		if i.Bad("integer", 1) {
			return
		}

		idx := i.Stack("integer").Pop().(int64)
		d.YankDup(idx)
	}

	ds.FunctionMap["stackdepth"] = func(d DataStack, i Interpreter) {
		i.Stack("integer").Push(d.Size())
	}
}

// Functions returns the FunctionMap of the datastack
func (s *datastack) Functions() FunctionMap {
	return s.FunctionMap
}

// Call calls a method from the FunctionMap
func (s *datastack) Call(method string, i Interpreter) {
	fn, ok := s.FunctionMap[method]
	if ok {
		fn(s, i)
	}
}

type NullDataStack struct {
	datastack
}

// Functions returns the FunctionMap of the datastack
func (s *NullDataStack) Functions() FunctionMap {
	return FunctionMap{}
}

// Call calls a method from the FunctionMap
func (s *NullDataStack) Call(method string, i Interpreter) {
	// Does nothing
}
