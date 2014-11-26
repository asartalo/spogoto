package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInterpreter(t *testing.T) {
	Convey("Given an interpreter", t, func() {
		i := NewInterpreter()
		d1 := tGenericDataStack(Elements{1, 2})
		i.RegisterStack("foo", d1)

		Convey("DataStacks are retrievable by name", func() {
			So(i.Stack("foo"), ShouldEqual, d1)
		})

		Convey("When Ok() is called", func() {
			Convey("On a non-existent stack should return false", func() {
				So(i.Ok("bar", 2), ShouldBeFalse)
			})

			Convey("On a stack that does not meet the number of elements should return false", func() {
				So(i.Ok("foo", 3), ShouldBeFalse)
			})

			Convey("On a stack that does meet the number of elements should return true", func() {
				So(i.Ok("foo", 2), ShouldBeTrue)
				So(i.Ok("foo", 1), ShouldBeTrue)
			})
		})

		Convey("When Bad() is called", func() {
			Convey("On a non-existent stack should return true", func() {
				So(i.Bad("bar", 2), ShouldBeTrue)
			})

			Convey("On a stack that does not meet the number of elements should return true", func() {
				So(i.Bad("foo", 3), ShouldBeTrue)
			})

			Convey("On a stack that does meet the number of elements should return false", func() {
				So(i.Bad("foo", 2), ShouldBeFalse)
				So(i.Bad("foo", 1), ShouldBeFalse)
			})
		})

		Convey("RandInt() should generate a random integer", func() {
			So(i.RandInt(), ShouldBeLessThan, 10)
			So(i.RandInt(), ShouldBeGreaterThan, -1)
		})

		Convey("RandFloat() should generate a random integer", func() {
			So(i.RandFloat(), ShouldBeLessThan, 1.0)
			So(i.RandFloat(), ShouldBeGreaterThanOrEqualTo, -0.0)
		})

		// Code
		Convey("When provided with code.", func() {
			code := "5 8 integer.+"

			Convey("And Run()", func() {
				result := i.Run(code)

				Convey("The code will be executed", func() {
					var value int64
					So(func() { value = i.Stack("integer").Pop().(int64) }, ShouldNotPanic)
					So(value, ShouldEqual, 13)
				})

				Convey("The result has the number of instructions executed", func() {
					So(result.InstructionCount, ShouldEqual, 3)
				})
			})
		})

	})

}

func TestCursorFunctions(t *testing.T) {
	testData := spogotoCodeTestData{
		{
			"Skipif with true",
			"1 true c.skipif 2 3",
			"Should skip next instruction",
			[]int64{}, []int64{1, 3},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Skipif with false",
			"1 false c.skipif 2 3",
			"Will not skip next instruction",
			[]int64{}, []int64{1, 2, 3},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Skipif with empty boolean",
			"1 c.skipif 2 3",
			"Will do nothing",
			[]int64{}, []int64{1, 2, 3},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"End",
			"1 c.end 2 3 4",
			"Will terminate code and ignore subsequent instructions",
			[]int64{}, []int64{1},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Endif with true",
			"1 true c.endif 2 3 4",
			"Will terminate code and ignore subsequent instructions",
			[]int64{}, []int64{1},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Endif with false",
			"1 false c.endif 2 3 4",
			"Will not terminate code",
			[]int64{}, []int64{1, 2, 3, 4},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Endif with empty boolean stack",
			"1 c.endif 2 3 4",
			"Will do nothing",
			[]int64{}, []int64{1, 2, 3, 4},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
	}
	for _, d := range testData {
		Convey(fmt.Sprintf("%s on code `%s`", d.toTest, d.code), t, func() {
			i := NewInterpreter()

			Convey(d.expectation, func() {
				So(func() { i.Run(d.code) }, ShouldNotPanic)
				So(i.Stack("integer").Elements(), ShouldResemble, int64Elements(d.intsAfter))
				So(i.Stack("boolean").Elements(), ShouldResemble, boolElements(d.boolsAfter))
				So(i.Stack("float").Elements(), ShouldResemble, float64Elements(d.floatsAfter))
			})
		})
	}
}
