package spogoto

import (
	"strconv"
)

func NewBooleanStack(bools []bool) *datastack {
	elements := Elements{}
	for _, v := range bools {
		elements = append(elements, bool(v))
	}
	d := NewDataStack(elements, FunctionMap{}, func(str string) (Element, bool) {
		val, err := strconv.ParseBool(str)
		return Element(val), err == nil
	})
	addBooleanFunctions(d)
	return d
}

func addBooleanFunctions(ds *datastack) {

	ds.FunctionMap["="] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		d.Push(d.Pop().(bool) == d.Pop().(bool))
	}

	ds.FunctionMap["and"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		b1 := d.Pop().(bool)
		b2 := d.Pop().(bool)
		d.Push(b1 && b2)
	}

	ds.FunctionMap["or"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		b1 := d.Pop().(bool)
		b2 := d.Pop().(bool)
		d.Push(b1 || b2)
	}

	ds.FunctionMap["xor"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		b1 := d.Pop().(bool)
		b2 := d.Pop().(bool)
		d.Push(b1 != b2)
	}

	ds.FunctionMap["not"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(1) {
			return
		}

		d.Push(!d.Pop().(bool))
	}

	ds.FunctionMap["frominteger"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("integer", 1) {
			return
		}

		d.Push(r.Stack("integer").Pop().(int64) != 0)
	}

	ds.FunctionMap["fromfloat"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("float", 1) {
			return
		}

		d.Push(r.Stack("float").Pop().(float64) != 0.0)
	}

}
