package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func int64Elements(ints []int64) Elements {
	elements := Elements{}
	for _, v := range ints {
		elements = append(elements, int64(v))
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

func TestIntegerStackArithmeticFunctions(t *testing.T) {
	i := NewInterpreter()

	testData := []struct {
		fn     string
		before []int64
		after  []int64
	}{
		{"+", []int64{1, 2, 6, 2}, []int64{1, 2, 8}},
		{"*", []int64{1, 2, 6, 2}, []int64{1, 2, 12}},
		{"-", []int64{1, 2, 6, 2}, []int64{1, 2, 4}},
		{"/", []int64{1, 2, 6, 2}, []int64{1, 2, 3}},
		{"/", []int64{6, 0}, []int64{6, 0}},
		{"%", []int64{1, 2, 6, 2}, []int64{1, 2, 0}},
		{"%", []int64{6, 0}, []int64{6, 0}},
		{"min", []int64{1, 2, 6, 2}, []int64{1, 2, 2}},
		{"max", []int64{1, 2, 6, 2}, []int64{1, 2, 6}},
	}

	for _, d := range testData {
		Convey(fmt.Sprintf("Call()ing '%s' should result %v", d.fn, d.after), t, func() {
			s := NewIntegerStack(d.before)
			So(func() { s.Call(d.fn, i) }, ShouldNotPanic)
			So(s.elements, ShouldResemble, int64Elements(d.after))
		})

	}
}

func TestIntegerStackBooleanFunctions(t *testing.T) {

	testData := []struct {
		fn     string
		before []int64
		after  []int64
		bools  []bool
	}{
		{
			">",
			[]int64{1, 2, 6, 2},
			[]int64{1, 2},
			[]bool{true},
		},
		{
			">",
			[]int64{1, 0, 1, 2},
			[]int64{1, 0},
			[]bool{false},
		},
		{
			"<",
			[]int64{1, 2, 6, 2},
			[]int64{1, 2},
			[]bool{false},
		},
		{
			"<",
			[]int64{1, 0, 1, 2},
			[]int64{1, 0},
			[]bool{true},
		},
		{
			"=",
			[]int64{9, 8, 7, 7},
			[]int64{9, 8},
			[]bool{true},
		},
		{
			"=",
			[]int64{9, 8, 7, 8},
			[]int64{9, 8},
			[]bool{false},
		},
	}

	for _, d := range testData {
		i := NewInterpreter()
		boolStack := NewDataStack(Elements{}, FunctionMap{})
		i.RegisterStack("boolean", boolStack)

		Convey(fmt.Sprintf("Call()ing '%s' with integer stack %v should result %v", d.fn, d.before, d.bools), t, func() {
			s := NewIntegerStack(d.before)
			s.Call(d.fn, i)
			So(s.elements, ShouldResemble, int64Elements(d.after))
			So(boolStack.elements, ShouldResemble, boolElements(d.bools))
		})

	}
}

func TestEmptyIntegerStack(t *testing.T) {
	i := NewInterpreter()
	elements := []int64{}

	functions := []string{
		"+", "*", "-",
		"/", "%", "min", "max",
		"<", "=", ">",
	}

	for _, fn := range functions {
		boolStack := NewDataStack(Elements{}, FunctionMap{})
		i.RegisterStack("boolean", boolStack)

		Convey(fmt.Sprintf("For empty stacks, Call()ing '%s' should not panic and do nothing", fn), t, func() {
			s := NewIntegerStack(elements)
			So(func() { s.Call(fn, i) }, ShouldNotPanic)
			So(s.elements, ShouldResemble, int64Elements([]int64{}))
			So(boolStack.elements, ShouldResemble, boolElements([]bool{}))
		})
	}
}
