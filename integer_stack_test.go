package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

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

	testData := spogotoTestData{
		{
			">",
			[]int64{1, 2, 6, 2}, []int64{1, 2},
			[]bool{}, []bool{true},
			[]float64{}, []float64{},
		},
		{
			">",
			[]int64{1, 0, 1, 2}, []int64{1, 0},
			[]bool{}, []bool{false},
			[]float64{}, []float64{},
		},
		{
			"<",
			[]int64{1, 2, 6, 2}, []int64{1, 2},
			[]bool{}, []bool{false},
			[]float64{}, []float64{},
		},
		{
			"<",
			[]int64{1, 0, 1, 2}, []int64{1, 0},
			[]bool{}, []bool{true},
			[]float64{}, []float64{},
		},
		{
			"=",
			[]int64{9, 8, 7, 7}, []int64{9, 8},
			[]bool{}, []bool{true},
			[]float64{}, []float64{},
		},
		{
			"=",
			[]int64{9, 8, 7, 8}, []int64{9, 8},
			[]bool{}, []bool{false},
			[]float64{}, []float64{},
		},
		{
			"fromboolean",
			[]int64{}, []int64{1},
			[]bool{true}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"fromboolean",
			[]int64{}, []int64{0},
			[]bool{false}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"fromfloat",
			[]int64{}, []int64{7},
			[]bool{}, []bool{},
			[]float64{7.35}, []float64{},
		},
	}

	for _, d := range testData {
		i := NewInterpreter()
		boolStack := NewDataStack(boolElements(d.boolsBefore), FunctionMap{})
		floatStack := NewDataStack(float64Elements(d.floatsBefore), FunctionMap{})
		i.RegisterStack("boolean", boolStack)
		i.RegisterStack("float", floatStack)

		Convey(fmt.Sprintf("Call()ing '%s' with integer stack %v should result %v", d.fn, d.intsBefore, d.boolsAfter), t, func() {
			s := NewIntegerStack(d.intsBefore)
			s.Call(d.fn, i)
			So(s.elements, ShouldResemble, int64Elements(d.intsAfter))
			So(boolStack.elements, ShouldResemble, boolElements(d.boolsAfter))
			So(floatStack.elements, ShouldResemble, float64Elements(d.floatsAfter))
		})

	}
}

func TestRandomInteger(t *testing.T) {
	Convey("Given an empty datastack", t, func() {
		i := NewInterpreter()
		s := NewIntegerStack([]int64{})

		Convey("When 'rand' is Call()ed", func() {
			s.Call("rand", i)

			Convey("It should generate a random integer from 0 to 9", func() {
				val := s.Pop()
				So(val, ShouldBeLessThan, 10)
				So(val, ShouldBeGreaterThan, -1)
			})
		})

	})
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
		floatStack := NewDataStack(Elements{}, FunctionMap{})
		i.RegisterStack("boolean", boolStack)
		i.RegisterStack("float", floatStack)

		Convey(fmt.Sprintf("For empty stacks, Call()ing '%s' should not panic and do nothing", fn), t, func() {
			s := NewIntegerStack(elements)
			So(func() { s.Call(fn, i) }, ShouldNotPanic)
			So(s.elements, ShouldResemble, int64Elements([]int64{}))
			So(boolStack.elements, ShouldResemble, boolElements([]bool{}))
			So(floatStack.elements, ShouldResemble, float64Elements([]float64{}))
		})
	}
}
