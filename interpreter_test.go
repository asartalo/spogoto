package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInterpreter(t *testing.T) {
	Convey("Given an Interpreter and a RunSet", t, func() {
		i := NewInterpreter(DefaultOptions)
		r := NewRunSet(i)
		d1 := tGenericDataStack(Elements{1, 2})
		r.RegisterStack("foo", d1)

		Convey("DataStacks are retrievable by name", func() {
			So(r.Stack("foo"), ShouldEqual, d1)
		})

		Convey("When Ok() is called", func() {
			Convey("On a non-existent stack should return false", func() {
				So(r.Ok("bar", 2), ShouldBeFalse)
			})

			Convey("On a stack that does not meet the number of elements should return false", func() {
				So(r.Ok("foo", 3), ShouldBeFalse)
			})

			Convey("On a stack that does meet the number of elements should return true", func() {
				So(r.Ok("foo", 2), ShouldBeTrue)
				So(r.Ok("foo", 1), ShouldBeTrue)
			})
		})

		Convey("When Bad() is called", func() {
			Convey("On a non-existent stack should return true", func() {
				So(r.Bad("bar", 2), ShouldBeTrue)
			})

			Convey("On a stack that does not meet the number of elements should return true", func() {
				So(r.Bad("foo", 3), ShouldBeTrue)
			})

			Convey("On a stack that does meet the number of elements should return false", func() {
				So(r.Bad("foo", 2), ShouldBeFalse)
				So(r.Bad("foo", 1), ShouldBeFalse)
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
					So(func() { value = result.Stack("integer").Pop().(int64) }, ShouldNotPanic)
					So(value, ShouldEqual, 13)
				})

				Convey("The result has the number of instructions executed", func() {
					So(result.InstructionCount(), ShouldEqual, 3)
				})
			})
		})

	})

}

func TestCursorFunctions(t *testing.T) {
	testData := spogotoCodeTestData{
		{
			"Skipif with true",
			"1 true cursor.skipif 2 3",
			"Should skip next instruction",
			[]int64{}, []int64{1, 3},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Skipif with false",
			"1 false cursor.skipif 2 3",
			"Will not skip next instruction",
			[]int64{}, []int64{1, 2, 3},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Skipif with empty boolean",
			"1 cursor.skipif 2 3",
			"Will do nothing",
			[]int64{}, []int64{1, 2, 3},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"End",
			"1 cursor.end 2 3 4",
			"Will terminate code and ignore subsequent instructions",
			[]int64{}, []int64{1},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Endif with true",
			"1 true cursor.endif 2 3 4",
			"Will terminate code and ignore subsequent instructions",
			[]int64{}, []int64{1},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Endif with false",
			"1 false cursor.endif 2 3 4",
			"Will not terminate code",
			[]int64{}, []int64{1, 2, 3, 4},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Endif with empty boolean stack",
			"1 cursor.endif 2 3 4",
			"Will do nothing",
			[]int64{}, []int64{1, 2, 3, 4},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Goto with value on integer stack",
			"3 cursor.goto 2 1 8",
			"Will go to cursor based on value from integer stack",
			[]int64{}, []int64{1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Goto with no value on integer stack",
			"3.0 cursor.goto 2 1 8",
			"Will do nothing",
			[]int64{}, []int64{2, 1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{3.0},
		},
		{
			"Goto with negative integer",
			"-3 cursor.goto 2 1 8",
			"Will do nothing",
			[]int64{}, []int64{2, 1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Goto with integer greater than code length",
			"10 cursor.goto 2 1 8",
			"Will do nothing",
			[]int64{}, []int64{2, 1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Gotoif with boolean true",
			"true 4 cursor.gotoif 2 1 8",
			"Will go to cursor based on value from integer stack",
			[]int64{}, []int64{1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Gotoif with boolean false",
			"false 4 cursor.gotoif 2 1 8",
			"Will go to cursor based on value from integer stack",
			[]int64{}, []int64{4, 2, 1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{},
		},
		{
			"Gotoif with nothing on boolean stack",
			"3.0 4 cursor.gotoif 2 1 8",
			"Will do nothing",
			[]int64{}, []int64{4, 2, 1, 8},
			[]bool{}, []bool{},
			[]float64{}, []float64{3.0},
		},
	}
	for _, d := range testData {
		Convey(fmt.Sprintf("%s on code `%s`", d.toTest, d.code), t, func() {
			i := NewInterpreter(DefaultOptions)
			var r RunSet

			Convey(d.expectation, func() {
				So(func() { r = i.Run(d.code) }, ShouldNotPanic)
				So(r.Stack("integer").Elements(), ShouldResemble, int64Elements(d.intsAfter))
				So(r.Stack("boolean").Elements(), ShouldResemble, boolElements(d.boolsAfter))
				So(r.Stack("float").Elements(), ShouldResemble, float64Elements(d.floatsAfter))
			})
		})
	}
}
