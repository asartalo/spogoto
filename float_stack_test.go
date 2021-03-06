package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestFloatStackFunctions(t *testing.T) {

	testData := spogotoTestData{
		floatsOnly("+", []float64{1.0, 6.0, 2.0}, []float64{1.0, 8.0}),
		floatsOnly("-", []float64{1.0, 6.0, 2.0}, []float64{1.0, 4.0}),
		floatsOnly("*", []float64{1.0, 6.0, 2.0}, []float64{1.0, 12.0}),
		floatsOnly("/", []float64{1.0, 6.0, 2.0}, []float64{1.0, 3.0}),
		floatsOnly("/", []float64{1.0, 6.0, 0.0}, []float64{1.0, 6.0, 0.0}),
		floatsOnly("%", []float64{1.0, 6.0, 2.0}, []float64{1.0, 0.0}),
		floatsOnly("%", []float64{1.0, 6.0, 0.0}, []float64{1.0, 6.0, 0.0}),
		floatsOnly("min", []float64{1.0, 6.0, 2.0}, []float64{1.0, 2.0}),
		floatsOnly("min", []float64{1.0, 6.0, 8.0}, []float64{1.0, 6.0}),
		floatsOnly("max", []float64{1.0, 6.0, 2.0}, []float64{1.0, 6.0}),
		floatsOnly("max", []float64{1.0, 6.0, 8.0}, []float64{1.0, 8.0}),
		floatsOnly("sin", []float64{7.0}, []float64{math.Sin(7.0)}),
		floatsOnly("cos", []float64{7.0}, []float64{math.Cos(7.0)}),
		floatsOnly("tan", []float64{7.0}, []float64{math.Tan(7.0)}),

		{
			">",
			[]int64{}, []int64{},
			[]bool{}, []bool{true},
			[]float64{1.0, 6.0, 2.0}, []float64{1.0},
		},
		{
			">",
			[]int64{}, []int64{},
			[]bool{}, []bool{false},
			[]float64{1.0, 2.0, 6.0}, []float64{1.0},
		},
		{
			"<",
			[]int64{}, []int64{},
			[]bool{}, []bool{false},
			[]float64{1.0, 6.0, 2.0}, []float64{1.0},
		},
		{
			"<",
			[]int64{}, []int64{},
			[]bool{}, []bool{true},
			[]float64{1.0, 2.0, 6.0}, []float64{1.0},
		},
		{
			"=",
			[]int64{}, []int64{},
			[]bool{}, []bool{true},
			[]float64{1.0, 2.0, 2.0}, []float64{1.0},
		},
		{
			"=",
			[]int64{}, []int64{},
			[]bool{}, []bool{false},
			[]float64{1.0, 2.0, 6.0}, []float64{1.0},
		},
		{
			"fromboolean",
			[]int64{}, []int64{},
			[]bool{true}, []bool{},
			[]float64{}, []float64{1.0},
		},
		{
			"fromboolean",
			[]int64{}, []int64{},
			[]bool{false}, []bool{},
			[]float64{}, []float64{0.0},
		},
		{
			"frominteger",
			[]int64{7}, []int64{},
			[]bool{}, []bool{},
			[]float64{}, []float64{7.0},
		},

		// Empties
		tnothingHappens("+"),
		tnothingHappens("-"),
		tnothingHappens("*"),
		tnothingHappens("/"),
		tnothingHappens("%"),
		tnothingHappens("min"),
		tnothingHappens("max"),
		tnothingHappens(">"),
		tnothingHappens("<"),
		tnothingHappens("="),
		tnothingHappens("sin"),
		tnothingHappens("cos"),
		tnothingHappens("tan"),
		tnothingHappens("fromboolean"),
		tnothingHappens("frominteger"),
	}

	for _, d := range testData {
		Convey(tPrimaryMessage("boolean", d), t, func() {
			i := NewInterpreter(DefaultOptions)
			r := NewRunSet(i)
			boolStack := tGenericDataStack(boolElements(d.boolsBefore))
			integerStack := tGenericDataStack(int64Elements(d.intsBefore))
			r.RegisterStack("boolean", boolStack)
			r.RegisterStack("integer", integerStack)

			s := NewFloatStack(d.floatsBefore)
			Convey("It shouldn't panic", func() {
				So(func() { s.Call(d.fn, r, i) }, ShouldNotPanic)

				Convey(fmt.Sprintf("Float elements should be %v", d.floatsAfter), func() {
					So(s.elements, ShouldResemble, float64Elements(d.floatsAfter))
				})

				Convey(tBoolMessaging(d.boolsBefore, d.boolsAfter), func() {
					So(boolStack.elements, ShouldResemble, boolElements(d.boolsAfter))
				})
				Convey(tIntMessaging(d.intsBefore, d.intsAfter), func() {
					So(integerStack.elements, ShouldResemble, int64Elements(d.intsAfter))
				})
			})
		})
	}
}

func TestOtherFloatStackFeatures(t *testing.T) {
	Convey("Given an empty float datastack", t, func() {
		i := NewInterpreter(DefaultOptions)
		r := NewRunSet(i)
		s := NewFloatStack([]float64{})

		Convey("When 'rand' is Call()ed", func() {
			s.Call("rand", r, i)

			Convey("It should generate a random float number", func() {
				val := s.Pop()
				So(val, ShouldBeLessThan, 1.0)
				So(val, ShouldBeGreaterThanOrEqualTo, 0)
			})
		})

		Convey("When pushing literals", func() {
			Convey("It should push proper float literal", func() {
				s.PushLiteral("2.0")
				So(s.Pop(), ShouldEqual, 2.0)
			})

			Convey("It should not push improper float literal", func() {
				s.PushLiteral("rrr")
				So(s.Pop(), ShouldBeNil)
			})
		})

	})
}

func TestFloatStackConstructor(t *testing.T) {
	Convey("Given an float stack constructor", t, func() {
		constructor := FloatStackConstructor

		Convey("When called", func() {
			stackType, stack := constructor()

			Convey("It should return float stack type", func() {
				So(stackType, ShouldEqual, "float")
			})

			// TODO: There must be a correct way to do this
			Convey("It should return a float stack", func() {
				So(stack.Elements(), ShouldResemble, NewFloatStack([]float64{}).Elements())
			})
		})
	})
}
