package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIntegerStackFunctions(t *testing.T) {
	testData := spogotoTestData{
		tintsOnly("+", []int64{1, 2, 6, 2}, []int64{1, 2, 8}),
		tintsOnly("*", []int64{1, 2, 6, 2}, []int64{1, 2, 12}),
		tintsOnly("-", []int64{1, 2, 6, 2}, []int64{1, 2, 4}),
		tintsOnly("/", []int64{1, 2, 6, 2}, []int64{1, 2, 3}),
		tintsOnly("/", []int64{6, 0}, []int64{6, 0}),
		tintsOnly("%", []int64{1, 2, 6, 2}, []int64{1, 2, 0}),
		tintsOnly("%", []int64{6, 0}, []int64{6, 0}),
		tintsOnly("min", []int64{1, 2, 6, 2}, []int64{1, 2, 2}),
		tintsOnly("min", []int64{1, 2, 6, 8}, []int64{1, 2, 6}),
		tintsOnly("max", []int64{1, 2, 6, 2}, []int64{1, 2, 6}),
		tintsOnly("max", []int64{1, 2, 6, 8}, []int64{1, 2, 8}),

		// Boolean Functions
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

		// Conversion functions
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

		// Empties
		tnothingHappens("+"),
		tnothingHappens("*"),
		tnothingHappens("-"),
		tnothingHappens("/"),
		tnothingHappens("%"),
		tnothingHappens("<"),
		tnothingHappens("="),
		tnothingHappens(">"),
		tnothingHappens("min"),
		tnothingHappens("max"),
		tnothingHappens("fromboolean"),
		tnothingHappens("fromfloat"),
	}

	for _, d := range testData {
		Convey(tPrimaryMessage("integer", d), t, func() {
			i := NewInterpreter()
			boolStack := NewDataStack(boolElements(d.boolsBefore), FunctionMap{})
			floatStack := NewDataStack(float64Elements(d.floatsBefore), FunctionMap{})
			i.RegisterStack("boolean", boolStack)
			i.RegisterStack("float", floatStack)

			s := NewIntegerStack(d.intsBefore)
			Convey("It shouldn't panic", func() {
				So(func() { s.Call(d.fn, i) }, ShouldNotPanic)

				Convey(fmt.Sprintf("Integer elements should be %v", d.intsAfter), func() {
					So(s.elements, ShouldResemble, int64Elements(d.intsAfter))
				})

				Convey(tBoolMessaging(d.boolsBefore, d.boolsAfter), func() {
					So(boolStack.elements, ShouldResemble, boolElements(d.boolsAfter))
				})
				Convey(tFloatMessaging(d.floatsBefore, d.floatsAfter), func() {
					So(floatStack.elements, ShouldResemble, float64Elements(d.floatsAfter))
				})

			})

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
