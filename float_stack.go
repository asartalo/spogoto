package spogoto

func NewFloatStack(floats []float64) *datastack {
	elements := Elements{}
	for _, v := range floats {
		elements = append(elements, float64(v))
	}
	d := &datastack{stack{elements}, FunctionMap{}}
	addFloatFunctions(d)
	return d
}

func addFloatFunctions(ds *datastack) {

	ds.FunctionMap[">"] = func(d DataStack, i Interpreter) {
		if d.Lack(2) {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		i.Stack("boolean").Push(f2 > f1)
	}

	ds.FunctionMap["<"] = func(d DataStack, i Interpreter) {
		if d.Lack(2) {
			return
		}

		f1 := d.Pop().(float64)
		f2 := d.Pop().(float64)

		i.Stack("boolean").Push(f2 < f1)
	}

}
