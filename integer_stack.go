package spogoto

import (
	"strconv"
)

func NewIntegerStack(ints []int64) *datastack {
	elements := Elements{}
	for _, v := range ints {
		elements = append(elements, int64(v))
	}

	d := NewDataStack(elements, FunctionMap{}, func(str string) (Element, bool) {
		val, err := strconv.ParseInt(str, 10, 64)
		return Element(val), err == nil
	})
	addIntegerFunctions(d)
	return d
}

func addIntegerFunctions(ds *datastack) {
	ds.FunctionMap["+"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		s := d.Pop().(int64) + d.Pop().(int64)
		d.Push(s)
	}

	ds.FunctionMap["*"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		s := d.Pop().(int64) * d.Pop().(int64)
		d.Push(s)
	}

	ds.FunctionMap["-"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		d.Push(-d.Pop().(int64) + d.Pop().(int64))
	}

	ds.FunctionMap["/"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) || d.Peek().(int64) == 0 {
			return
		}

		i1 := d.Pop().(int64)
		i2 := d.Pop().(int64)

		d.Push(i2 / i1)
	}

	ds.FunctionMap["%"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) || d.Peek().(int64) == 0 {
			return
		}

		i1 := d.Pop().(int64)
		i2 := d.Pop().(int64)

		d.Push(i2 % i1)
	}

	ds.FunctionMap["min"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		i1 := d.Pop().(int64)
		i2 := d.Pop().(int64)

		if i1 < i2 {
			d.Push(i1)
		} else {
			d.Push(i2)
		}
	}

	ds.FunctionMap["max"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		i1 := d.Pop().(int64)
		i2 := d.Pop().(int64)

		if i1 > i2 {
			d.Push(i1)
		} else {
			d.Push(i2)
		}
	}

	ds.FunctionMap[">"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		i1 := d.Pop().(int64)
		i2 := d.Pop().(int64)

		r.Stack("boolean").Push(i2 > i1)
	}

	ds.FunctionMap["<"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		i1 := d.Pop().(int64)
		i2 := d.Pop().(int64)

		r.Stack("boolean").Push(i2 < i1)
	}

	ds.FunctionMap["="] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		r.Stack("boolean").Push(d.Pop().(int64) == d.Pop().(int64))
	}

	ds.FunctionMap["fromboolean"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("boolean", 1) {
			return
		}

		b := r.Stack("boolean").Pop().(bool)
		if b {
			d.Push(int64(1))
		} else {
			d.Push(int64(0))
		}
	}

	ds.FunctionMap["fromfloat"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("float", 1) {
			return
		}

		f := r.Stack("float").Pop().(float64)
		d.Push(int64(f))
	}

	ds.FunctionMap["rand"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Push(i.RandInt())
	}

}
