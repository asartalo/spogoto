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

func tBoolPrimaryMessaging(d spogotoTestItem) (message string) {
	message = fmt.Sprintf("Call()ing '%s' with boolean stack %v", d.fn, d.boolsBefore)
	if len(d.intsBefore) > 0 {
		message += fmt.Sprintf(" with integer stack %v", d.intsBefore)
	}
	if len(d.floatsBefore) > 0 {
		message += fmt.Sprintf(" with float stack %v", d.floatsBefore)
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
