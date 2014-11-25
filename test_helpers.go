package spogoto

import (
	"fmt"
)

type spogotoTestItem struct {
	fn           string
	intsBefore   []int64
	intsAfter    []int64
	boolsBefore  []bool
	boolsAfter   []bool
	floatsBefore []float64
	floatsAfter  []float64
}

type spogotoTestData []spogotoTestItem

type spogotoCodeTestItem struct {
	fn           string
	intsBefore   []int64
	intsAfter    []int64
	boolsBefore  []bool
	boolsAfter   []bool
	floatsBefore []float64
	floatsAfter  []float64
	code         string
}

type spogotoCodeTestData []spogotoCodeTestItem

func int64Elements(ints []int64) Elements {
	elements := Elements{}
	for _, v := range ints {
		elements = append(elements, int64(v))
	}
	return elements
}

func float64Elements(ints []float64) Elements {
	elements := Elements{}
	for _, v := range ints {
		elements = append(elements, float64(v))
	}
	return elements
}

func boolElements(bools []bool) Elements {
	elements := Elements{}
	for _, v := range bools {
		elements = append(elements, bool(v))
	}
	return elements
}

func tintsOnly(fn string, is1 []int64, is2 []int64) spogotoTestItem {
	return spogotoTestItem{
		fn,
		is1, is2,
		[]bool{}, []bool{},
		[]float64{}, []float64{},
	}

}

func floatsOnly(fn string, fs1 []float64, fs2 []float64) spogotoTestItem {
	return spogotoTestItem{
		fn,
		[]int64{}, []int64{},
		[]bool{}, []bool{},
		fs1, fs2,
	}
}

func tbooleansOnly(fn string, bs1 []bool, bs2 []bool) spogotoTestItem {
	return spogotoTestItem{
		fn,
		[]int64{}, []int64{},
		bs1, bs2,
		[]float64{}, []float64{},
	}
}

func tnothingHappens(fn string) spogotoTestItem {
	return tbooleansOnly(fn, []bool{}, []bool{})
}

func tPrimaryMessage(t string, d spogotoTestItem) (message string) {
	var before interface{}
	switch t {
	case "integer":
		before = d.intsBefore
	case "float":
		before = d.floatsBefore
	case "boolean":
		before = d.boolsBefore
	}

	message = fmt.Sprintf("Call()ing '%s' with %s stack %v", d.fn, t, before)
	if len(d.intsBefore) > 0 && t != "integer" {
		message += fmt.Sprintf(" with integer stack %v", d.intsBefore)
	}
	if len(d.boolsBefore) > 0 && t != "boolean" {
		message += fmt.Sprintf(" and boolean stack %v", d.boolsBefore)
	}
	if len(d.floatsBefore) > 0 && t != "float" {
		message += fmt.Sprintf(" and float stack %v", d.floatsBefore)
	}

	return message

}

func tBoolMessaging(before []bool, after []bool) (message string) {
	if len(before) == 0 && len(after) == 0 {
		message = "It shouldn't modify the boolean stack"
	} else {
		message = fmt.Sprintf(
			"It should modify the boolean stack from %v to %v",
			before, after,
		)
	}

	return message
}

func tIntMessaging(before []int64, after []int64) (message string) {
	if len(before) == 0 && len(after) == 0 {
		message = "It shouldn't modify the integer stack"
	} else {
		message = fmt.Sprintf(
			"It should modify the integer stack from %v to %v",
			before, after,
		)
	}

	return message
}

func tFloatMessaging(before []float64, after []float64) (message string) {
	if len(before) == 0 && len(after) == 0 {
		message = "It shouldn't modify the float stack"
	} else {
		message = fmt.Sprintf(
			"It should modify the float stack from %v to %v",
			before, after,
		)
	}

	return message
}

func tGenericDataStack(els Elements) *datastack {
	return NewDataStack(els, FunctionMap{}, func(strVal string) (Element, bool) { return Element(strVal), false })
}
