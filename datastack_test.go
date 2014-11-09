package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDataStack(t *testing.T) {
	Convey("Given a stack with 4 elements and a function definition", t, func() {
		elements := Elements{"a", "b", "c", "d"}
		functions := make(FunctionMap)
		i := Interpreter{}
		functions["foo"] = func(s DataStack, i Interpreter) {
			s.Push("e")
		}
		s := NewDataStack(elements, functions)

		Convey("Size works", func() {
			So(s.Size(), ShouldEqual, 4)
		})

		Convey("FunctionMap() should return function map", func() {
			So(s.Functions(), ShouldEqual, functions)
		})

		Convey("When a function is Call()ed", func() {
			s.Call("foo", i)

			Convey("It should execute the function", func() {
				So(s.Peek(), ShouldEqual, "e")
			})
		})

		Convey("It should not panic when a non-existent function is Call()ed", func() {
			So(func() { s.Call("bar", i) }, ShouldNotPanic)
		})

		Convey("When 'pop' is Call()ed", func() {
			s.Call("pop", i)

			Convey("It should remove the top element of the stack", func() {
				So(s.Peek(), ShouldEqual, "c")
				So(s.Size(), ShouldEqual, 3)
			})
		})

		Convey("When 'rotate' is Call()ed", func() {
			s.Call("rotate", i)

			Convey("It should rotate the top 3 values on the stack", func() {
				So(s.Pop(), ShouldEqual, "b")
				So(s.Pop(), ShouldEqual, "d")
				So(s.Pop(), ShouldEqual, "c")
			})
		})
	})

}
