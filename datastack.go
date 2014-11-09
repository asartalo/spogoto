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

func NewDataStack(elements Elements, functions FunctionMap) datastack {
	d := datastack{stack{elements}, functions}
	addCommonFunctions(d)
	return d
}

func addCommonFunctions(ds datastack) {

	ds.FunctionMap["pop"] = func(d DataStack, i Interpreter) {
		d.Pop()
	}

	ds.FunctionMap["rotate"] = func(d DataStack, i Interpreter) {
		d.Rotate()
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
