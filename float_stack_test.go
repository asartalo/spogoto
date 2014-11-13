package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestFloatStackFunctions(t *testing.T) {
	floatOnly := func(fn string, fs1 []float64, fs2 []float64) spogotoTestItem {
		return spogotoTestItem{
			fn,
			[]int64{}, []int64{},
			[]bool{}, []bool{},
			fs1, fs2,
		}
	}

	nothingHappens := func(fn string) spogotoTestItem {
		return spogotoTestItem{
			fn,
			[]int64{}, []int64{},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		}
	}

	testData := spogotoTestData{
		floatOnly("+", []float64{1.0, 6.0, 2.0}, []float64{1.0, 8.0}),
		floatOnly("-", []float64{1.0, 6.0, 2.0}, []float64{1.0, 4.0}),
		floatOnly("*", []float64{1.0, 6.0, 2.0}, []float64{1.0, 12.0}),
		floatOnly("/", []float64{1.0, 6.0, 2.0}, []float64{1.0, 3.0}),
		floatOnly("/", []float64{1.0, 6.0, 0.0}, []float64{1.0, 6.0, 0.0}),
		floatOnly("%", []float64{1.0, 6.0, 2.0}, []float64{1.0, 0.0}),
		floatOnly("%", []float64{1.0, 6.0, 0.0}, []float64{1.0, 6.0, 0.0}),
		floatOnly("min", []float64{1.0, 6.0, 2.0}, []float64{1.0, 2.0}),
		floatOnly("max", []float64{1.0, 6.0, 2.0}, []float64{1.0, 6.0}),
		floatOnly("sin", []float64{7.0}, []float64{math.Sin(7.0)}),
		floatOnly("cos", []float64{7.0}, []float64{math.Cos(7.0)}),
		floatOnly("tan", []float64{7.0}, []float64{math.Tan(7.0)}),

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
		nothingHappens("+"),
		nothingHappens("-"),
		nothingHappens("*"),
		nothingHappens("/"),
		nothingHappens("%"),
		nothingHappens("min"),
		nothingHappens("max"),
		nothingHappens(">"),
		nothingHappens("<"),
		nothingHappens("="),
		nothingHappens("sin"),
		nothingHappens("cos"),
		nothingHappens("tan"),
	}

	for _, d := range testData {
		i := NewInterpreter()
		boolStack := NewDataStack(boolElements(d.boolsBefore), FunctionMap{})
		integerStack := NewDataStack(int64Elements(d.intsBefore), FunctionMap{})
		i.RegisterStack("boolean", boolStack)
		i.RegisterStack("integer", integerStack)

		Convey(fmt.Sprintf("Call()ing '%s' with float stack %v", d.fn, d.floatsBefore), t, func() {
			s := NewFloatStack(d.floatsBefore)
			So(func() { s.Call(d.fn, i) }, ShouldNotPanic)
			So(s.elements, ShouldResemble, float64Elements(d.floatsAfter))
			So(boolStack.elements, ShouldResemble, boolElements(d.boolsAfter))
			So(integerStack.elements, ShouldResemble, int64Elements(d.intsAfter))
		})

	}
}

func TestRandomFloat(t *testing.T) {
	Convey("Given an empty float datastack", t, func() {
		i := NewInterpreter()
		s := NewFloatStack([]float64{})

		Convey("When 'rand' is Call()ed", func() {
			s.Call("rand", i)

			Convey("It should generate a random float number", func() {
				val := s.Pop()
				So(val, ShouldBeLessThan, 1.0)
				So(val, ShouldBeGreaterThanOrEqualTo, 0)
			})
		})

	})
}
