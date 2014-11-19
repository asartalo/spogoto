package spogoto

import (
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
			code := "1 2 integer.+"

			Convey("And Run()", func() {
				i.Run(code)

				Convey("The code will be executed", func() {
					var result int64
					So(func() { result = i.Stack("integer").Pop().(int64) }, ShouldNotPanic)
					So(result, ShouldEqual, 3)
				})
			})
		})

	})

}
