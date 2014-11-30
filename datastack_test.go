package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDataStack(t *testing.T) {
	i := NewInterpreterDefault()
	r := NewRunSet(i)

	Convey("Given a stack with 4 elements and a function definition", t, func() {
		elements := Elements{"a", "b", "c", "d"}
		functions := make(FunctionMap)
		functions["foo"] = func(s DataStack, r RunSet, i Interpreter) {
			s.Push("e")
		}
		cfunc := func(strVal string) (Element, bool) { return Element(strVal), strVal == "x" }
		s := NewDataStack(elements, functions, cfunc)

		Convey("FunctionMap() should return function map", func() {
			So(s.Functions(), ShouldEqual, functions)
		})

		Convey("When PushLiteral() is called", func() {
			s.PushLiteral("x")
			s.PushLiteral("z")

			Convey("Uses ConversionFunc to push acceptable values", func() {
				So(s.Pop(), ShouldEqual, "x")
			})
		})

		Convey("When a function is Call()ed", func() {
			s.Call("foo", r, i)

			Convey("It should execute the function", func() {
				So(s.Peek(), ShouldEqual, "e")
			})
		})

		Convey("It should not panic when a non-existent function is Call()ed", func() {
			So(func() { s.Call("bar", r, i) }, ShouldNotPanic)
		})

		Convey("When 'pop' is Call()ed", func() {
			s.Call("pop", r, i)

			Convey("It should remove the top element of the stack", func() {
				So(s.Peek(), ShouldEqual, "c")
				So(s.Size(), ShouldEqual, 3)
			})
		})

		Convey("When 'swap' is Call()ed", func() {
			s.Call("swap", r, i)
			Convey("It should swap the positions of the top two elements", func() {
				So(s.Pop(), ShouldEqual, "c")
				So(s.Pop(), ShouldEqual, "d")
			})
		})

		Convey("When 'dup' is Call()ed", func() {
			s.Call("dup", r, i)
			Convey("It should duplicate the top element", func() {
				So(s.Size(), ShouldEqual, 5)
				So(s.Pop(), ShouldEqual, "d")
				So(s.Pop(), ShouldEqual, "d")
			})
		})

		Convey("When 'rotate' is Call()ed", func() {
			s.Call("rotate", r, i)

			Convey("It should rotate the top 3 values on the stack", func() {
				So(s.Pop(), ShouldEqual, "b")
				So(s.Pop(), ShouldEqual, "d")
				So(s.Pop(), ShouldEqual, "c")
			})
		})

		Convey("When 'flush' is Call()ed", func() {
			s.Call("flush", r, i)

			Convey("It should empty the stack", func() {
				So(s.Size(), ShouldEqual, 0)
			})
		})

		Convey("And a registered integer stack with element 2", func() {
			iStack := tGenericDataStack(Elements{int64(2)})
			r.RegisterStack("integer", iStack)

			Convey("When 'shove' is Call()ed", func() {
				s.Call("shove", r, i)

				Convey("It should shove the top most element to the 2 position", func() {
					So(s.Pop(), ShouldEqual, "c")
					So(s.Pop(), ShouldEqual, "b")
					So(s.Pop(), ShouldEqual, "d")
				})
			})

			Convey("When 'yank' is Call()ed", func() {
				s.Call("yank", r, i)

				Convey("It pulls the 2nd element off the stack and places it on top", func() {
					So(s.Pop(), ShouldEqual, "b")
				})
			})

			Convey("When 'yankdup' is Call()ed", func() {
				s.Call("yankdup", r, i)

				Convey("It copies the 2nd element stack and places the copy on top", func() {
					So(s.Size(), ShouldEqual, 5)
					So(s.Pop(), ShouldEqual, "b")
				})
			})

			Convey("When 'stackdepth' is Call()ed", func() {
				s.Call("stackdepth", r, i)

				Convey("It should push the size of the stack to the integer stack", func() {
					So(iStack.Pop(), ShouldEqual, 4)
				})
			})

		})
	})

	Convey("Given an empty datastack", t, func() {
		s := tGenericDataStack(Elements{})

		Convey("And a registered integer stack with no elements", func() {
			iStack := tGenericDataStack(Elements{})
			r.RegisterStack("integer", iStack)

			Convey("A Call() to 'shove' must not panic", func() {
				So(func() { s.Call("shove", r, i) }, ShouldNotPanic)
			})

			Convey("A Call() to 'yank' must not panic", func() {
				So(func() { s.Call("yank", r, i) }, ShouldNotPanic)
			})

			Convey("A Call() to 'yankdup' must not panic", func() {
				So(func() { s.Call("yankdup", r, i) }, ShouldNotPanic)
			})

		})
	})

}
