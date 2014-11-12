package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInterpreter(t *testing.T) {
	Convey("Given an interpreter", t, func() {
		i := NewInterpreter()
		d1 := NewDataStack(Elements{1, 2}, FunctionMap{})
		// d2 := NewDataStack(Elements{}, FunctionMap{})
		i.RegisterStack("foo", d1)
		// i.RegisterStack("bar", d2)

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
	})
}
