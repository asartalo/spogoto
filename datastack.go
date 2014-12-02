package spogoto

// FunctionMap is a map of functions that operate on the DataStack and
// other DataStacks accessible through the RunSet.
type FunctionMap map[string]func(DataStack, RunSet, Interpreter)

// DataStack is a Stack used by the Interpreter
// to store data of a specific type and has functions that can manipulate
// data from the interpreter
type DataStack interface {
	Stack
	Functions() FunctionMap
	Call(string, RunSet, Interpreter)
	PushLiteral(string)
}

// ConversionFunc is a function that converts a string literal to an element
// of an appropriate type.
type ConversionFunc func(string) (Element, bool)

type datastack struct {
	stack
	FunctionMap    FunctionMap
	ConversionFunc ConversionFunc
}

// NewDataStack is a constructor for datastack.
func NewDataStack(elements Elements, functions FunctionMap, fn ConversionFunc) *datastack {
	d := &datastack{stack{elements}, functions, fn}
	addCommonFunctions(d)
	return d
}

func addCommonFunctions(ds *datastack) {

	ds.FunctionMap["pop"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Pop()
	}

	ds.FunctionMap["swap"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Swap()
	}

	ds.FunctionMap["rotate"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Rotate()
	}

	ds.FunctionMap["shove"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("integer", 1) {
			return
		}

		idx := r.Stack("integer").Pop().(int64)
		d.Shove(d.Pop(), idx)
	}

	ds.FunctionMap["yank"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("integer", 1) {
			return
		}

		idx := r.Stack("integer").Pop().(int64)
		d.Yank(idx)
	}

	ds.FunctionMap["yankdup"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("integer", 1) {
			return
		}

		idx := r.Stack("integer").Pop().(int64)
		d.YankDup(idx)
	}

	ds.FunctionMap["stackdepth"] = func(d DataStack, r RunSet, i Interpreter) {
		r.Stack("integer").Push(d.Size())
	}

	ds.FunctionMap["flush"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Flush()
	}

	ds.FunctionMap["dup"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Dup()
	}
}

// Functions returns the FunctionMap of the datastack
func (s *datastack) Functions() FunctionMap {
	return s.FunctionMap
}

// Call calls a method from the FunctionMap
func (s *datastack) Call(method string, r RunSet, i Interpreter) {
	fn, ok := s.FunctionMap[method]
	if ok {
		fn(s, r, i)
	}
}

// PushLiteral converts a string literal to an appropriate type
// and adds it to the stack.
func (s *datastack) PushLiteral(sval string) {
	el, ok := s.ConversionFunc(sval)
	if ok {
		s.Push(el)
	}
}

// A NullDataStack is a DataStack that has nothing and does nothing.
type NullDataStack struct {
	datastack
}

// Functions returns a FunctionMap that is always empty.
func (s *NullDataStack) Functions() FunctionMap {
	return FunctionMap{}
}

// Call calls a method from the FunctionMap which in NullDataStack's case is nothing.
func (s *NullDataStack) Call(method string, r RunSet, i Interpreter) {
	// Does nothing
}

// PushLiteral accepts a string literal but does practically nothing.
func (s *NullDataStack) PushLiteral(sval string) {
}
