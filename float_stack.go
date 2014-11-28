package spogoto

import (
	"math"
	"strconv"
)

func NewFloatStack(floats []float64) *datastack {
	elements := Elements{}
	for _, v := range floats {
		elements = append(elements, float64(v))
	}
	d := NewDataStack(elements, FunctionMap{}, func(str string) (Element, bool) {
		val, err := strconv.ParseFloat(str, 64)
		return Element(val), err == nil
	})
	addFloatFunctions(d)
	return d
}

func addFloatFunctions(ds *datastack) {

	ds.FunctionMap["+"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		d.Push(d.Pop().(float64) + d.Pop().(float64))
	}

	ds.FunctionMap["*"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		d.Push(d.Pop().(float64) * d.Pop().(float64))
	}

	ds.FunctionMap["-"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		d.Push(-d.Pop().(float64) + d.Pop().(float64))
	}

	ds.FunctionMap["/"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) || d.Peek().(float64) == 0 {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		d.Push(f2 / f1)
	}

	ds.FunctionMap["%"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) || d.Peek().(float64) == 0 {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		mod := math.Mod(f2, f1)
		d.Push(mod)
	}

	ds.FunctionMap["min"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		if f1 < f2 {
			d.Push(f1)
		} else {
			d.Push(f2)
		}
	}

	ds.FunctionMap["max"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		if f1 > f2 {
			d.Push(f1)
		} else {
			d.Push(f2)
		}
	}

	ds.FunctionMap[">"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		r.Stack("boolean").Push(f2 > f1)
	}

	ds.FunctionMap["<"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		r.Stack("boolean").Push(f2 < f1)
	}

	ds.FunctionMap["="] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(2) {
			return
		}

		r.Stack("boolean").Push(d.Pop().(float64) == d.Pop().(float64))
	}

	ds.FunctionMap["fromboolean"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("boolean", 1) {
			return
		}

		b := r.Stack("boolean").Pop().(bool)
		if b {
			d.Push(float64(1))
		} else {
			d.Push(float64(0))
		}
	}

	ds.FunctionMap["frominteger"] = func(d DataStack, r RunSet, i Interpreter) {
		if r.Bad("integer", 1) {
			return
		}

		d.Push(float64(r.Stack("integer").Pop().(int64)))
	}

	ds.FunctionMap["sin"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(1) {
			return
		}

		d.Push(math.Sin(d.Pop().(float64)))
	}

	ds.FunctionMap["cos"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(1) {
			return
		}

		d.Push(math.Cos(d.Pop().(float64)))
	}

	ds.FunctionMap["tan"] = func(d DataStack, r RunSet, i Interpreter) {
		if d.Lack(1) {
			return
		}

		d.Push(math.Tan(d.Pop().(float64)))
	}

	ds.FunctionMap["rand"] = func(d DataStack, r RunSet, i Interpreter) {
		d.Push(i.RandFloat())
	}

}
