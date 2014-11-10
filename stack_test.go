package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack(t *testing.T) {
	Convey("Given a stack with 4 elements", t, func() {
		elements := Elements{"a", "b", "c", "d"}
		s := NewStack(elements)

		Convey("It's size should be the number of elements", func() {
			So(s.Size(), ShouldEqual, 4)
		})

		Convey("It is not empty", func() {
			So(s.IsEmpty(), ShouldBeFalse)
		})

		Convey("When Peek() is called", func() {
			e := s.Peek()

			Convey("It should return the top element of the stack", func() {
				So(e, ShouldEqual, "d")
			})

			Convey("It should not change the number of elements on the stack", func() {
				So(s.Size(), ShouldEqual, 4)
			})
		})

		Convey("When the stack is Pop()ed", func() {
			e := s.Pop()

			Convey("It should return the top element of the stack", func() {
				So(e, ShouldEqual, "d")
			})

			Convey("It should remove the element from the stack", func() {
				So(s.Size(), ShouldEqual, 3)
			})
		})

		Convey("When an element is Push()ed", func() {
			s.Push("e")

			Convey("It should add the element to the top of the stack", func() {
				So(s.Peek(), ShouldEqual, "e")
				So(s.Size(), ShouldEqual, 5)
			})
		})

		Convey("When the stack is Swap()ed", func() {
			s.Swap()

			Convey("It should swap the top two elements on the stack", func() {
				So(s.Pop(), ShouldEqual, "c")
				So(s.Pop(), ShouldEqual, "d")
			})
		})

		Convey("When the stack is Flush()ed", func() {
			s.Flush()

			Convey("It should empty the stack", func() {
				So(s.Size(), ShouldEqual, 0)
			})
		})

		Convey("When the stack is Rotate()d", func() {
			s.Rotate()

			Convey("The top element should be the third before", func() {
				So(s.Pop(), ShouldEqual, "b")
				So(s.Pop(), ShouldEqual, "d")
				So(s.Pop(), ShouldEqual, "c")
			})
		})

		Convey("When an element is Yank()ed", func() {
			Convey("With 3", func() {
				s.Yank(3)

				Convey("That element should now be on top", func() {
					So(s.Pop(), ShouldEqual, "a")
					So(s.Pop(), ShouldEqual, "d")
					So(s.Pop(), ShouldEqual, "c")
				})

				Convey("No new elements are added", func() {
					So(s.Size(), ShouldEqual, 4)
				})
			})

			Convey("With 1", func() {
				s.Yank(1)

				Convey("That element should now be on top", func() {
					So(s.Pop(), ShouldEqual, "c")
					So(s.Pop(), ShouldEqual, "d")
					So(s.Pop(), ShouldEqual, "b")
				})
			})

			Convey("With -1", func() {
				s.Yank(-1)
				Convey("Should do nothing", func() {
					So(s.Pop(), ShouldEqual, "d")
				})
			})

			Convey("With a number greater than the last index", func() {
				s.Yank(4)
				Convey("Should do nothing", func() {
					So(s.Pop(), ShouldEqual, "d")
				})
			})

		})

		Convey("When Dup() is called", func() {
			s.Dup()
			Convey("It should copy the top element", func() {
				So(s.Pop(), ShouldEqual, "d")
				So(s.Pop(), ShouldEqual, "d")
			})
		})

		Convey("When an element is YankDup()ed", func() {
			Convey("With 3", func() {
				s.YankDup(1)

				Convey("That copy of the element should now be on top", func() {
					So(s.Pop(), ShouldEqual, "c")
					So(s.Pop(), ShouldEqual, "d")
					So(s.Pop(), ShouldEqual, "c")
				})

				Convey("The element is added", func() {
					So(s.Size(), ShouldEqual, 5)
				})
			})

			Convey("With -1", func() {
				s.Yank(-1)
				Convey("Should do nothing", func() {
					So(s.Pop(), ShouldEqual, "d")
				})
			})

			Convey("With a number greater than the last index", func() {
				s.Yank(4)
				Convey("Should do nothing", func() {
					So(s.Pop(), ShouldEqual, "d")
				})
			})
		})

		Convey("When an element is Shove()d", func() {
			s.Shove("x", 2)

			Convey("It should be placed on the specified index", func() {
				So(s.Pop(), ShouldEqual, "d")
				So(s.Pop(), ShouldEqual, "c")
				So(s.Pop(), ShouldEqual, "x")
			})
		})

		Convey("Has() with number of elements within Size() should return true", func() {
			So(s.Has(2), ShouldBeTrue)
			So(s.Has(4), ShouldBeTrue)
			So(s.Has(1), ShouldBeTrue)
			So(s.Has(5), ShouldBeFalse)
		})

		Convey("Lack() with number of elements within Size() should return false", func() {
			So(s.Lack(2), ShouldBeFalse)
			So(s.Lack(4), ShouldBeFalse)
			So(s.Lack(1), ShouldBeFalse)
			So(s.Lack(5), ShouldBeTrue)
		})

	})

	Convey("Given an empty stack", t, func() {
		s := NewStack(Elements{})

		Convey("It is empty", func() {
			So(s.IsEmpty(), ShouldBeTrue)
		})

		Convey("Peek() should return nil", func() {
			So(s.Peek(), ShouldBeNil)
		})

		Convey("Pop() should return nil", func() {
			So(s.Pop(), ShouldBeNil)
		})

		Convey("Swap() should do nothing and not panic", func() {
			So(func() { s.Swap() }, ShouldNotPanic)
		})

		Convey("Rotate() should do nothing and not panic", func() {
			So(func() { s.Rotate() }, ShouldNotPanic)
		})

		Convey("Yank() should do nothing and not panic", func() {
			So(func() { s.Yank(2) }, ShouldNotPanic)
		})

		Convey("YankDup() should do nothing and not panic", func() {
			So(func() { s.YankDup(1) }, ShouldNotPanic)
		})

		Convey("Dup() should do nothing and not panic", func() {
			So(func() { s.Dup() }, ShouldNotPanic)
		})

		Convey("Shove() should do nothing and not panic", func() {
			So(func() { s.Shove("x", 1) }, ShouldNotPanic)
		})
	})
}
