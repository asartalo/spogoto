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
			i := NewInterpreter(DefaultOptions)
			r := NewRunSet(i)
			boolStack := tGenericDataStack(boolElements(d.boolsBefore))
			floatStack := tGenericDataStack(float64Elements(d.floatsBefore))
			r.RegisterStack("boolean", boolStack)
			r.RegisterStack("float", floatStack)

			s := NewIntegerStack(d.intsBefore)
			Convey("It shouldn't panic", func() {
				So(func() { s.Call(d.fn, r, i) }, ShouldNotPanic)

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

func TestIntegerOtherFeatures(t *testing.T) {
	Convey("Given an empty integer stack", t, func() {
		i := NewInterpreter(DefaultOptions)
		r := NewRunSet(i)
		s := NewIntegerStack([]int64{})

		Convey("When 'rand' is Call()ed", func() {
			s.Call("rand", r, i)

			Convey("It should generate a random integer from 0 to 9", func() {
				val := s.Pop()
				So(val, ShouldBeLessThan, 10)
				So(val, ShouldBeGreaterThan, -1)
			})
		})

		Convey("When pushing literals", func() {
			Convey("It should push proper integer literal", func() {
				s.PushLiteral("2")
				So(s.Pop(), ShouldEqual, 2)
			})

			Convey("It should not push improper integer literal", func() {
				s.PushLiteral("rrr")
				So(s.Pop(), ShouldBeNil)
			})
		})

	})
}

func TestIntegerStackConstructor(t *testing.T) {
	Convey("Given an integer stack constructor", t, func() {
		constructor := IntegerStackConstructor

		Convey("When called", func() {
			stackType, stack := constructor()

			Convey("It should return integer stack type", func() {
				So(stackType, ShouldEqual, "integer")
			})

			Convey("It should return an integer stack", func() {
				So(stack.Elements(), ShouldResemble, NewIntegerStack([]int64{}).Elements())
			})
		})
	})
}
