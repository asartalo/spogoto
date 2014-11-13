package spogoto

func NewBooleanStack(bools []bool) *datastack {
	elements := Elements{}
	for _, v := range bools {
		elements = append(elements, bool(v))
	}
	d := &datastack{stack{elements}, FunctionMap{}}
	addBooleanFunctions(d)
	return d
}

func addBooleanFunctions(ds *datastack) {

	ds.FunctionMap["="] = func(d DataStack, i Interpreter) {
		if d.Lack(2) {
			return
		}

		d.Push(d.Pop().(bool) == d.Pop().(bool))
	}

	ds.FunctionMap["and"] = func(d DataStack, i Interpreter) {
		if d.Lack(2) {
			return
		}

		b1 := d.Pop().(bool)
		b2 := d.Pop().(bool)
		d.Push(b1 && b2)
	}

	ds.FunctionMap["or"] = func(d DataStack, i Interpreter) {
		if d.Lack(2) {
			return
		}

		b1 := d.Pop().(bool)
		b2 := d.Pop().(bool)
		d.Push(b1 || b2)
	}

	ds.FunctionMap["xor"] = func(d DataStack, i Interpreter) {
		if d.Lack(2) {
			return
		}

		b1 := d.Pop().(bool)
		b2 := d.Pop().(bool)
		d.Push(b1 != b2)
	}

	ds.FunctionMap["not"] = func(d DataStack, i Interpreter) {
		if d.Lack(1) {
			return
		}

		d.Push(!d.Pop().(bool))
	}

	ds.FunctionMap["frominteger"] = func(d DataStack, i Interpreter) {
		if i.Bad("integer", 1) {
			return
		}

		d.Push(i.Stack("integer").Pop().(int64) != 0)
	}

	ds.FunctionMap["fromfloat"] = func(d DataStack, i Interpreter) {
		if i.Bad("float", 1) {
			return
		}

		d.Push(i.Stack("float").Pop().(float64) != 0.0)
	}

}
