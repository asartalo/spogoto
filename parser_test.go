package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParser(t *testing.T) {
	Convey("Given a parser", t, func() {
		parser := NewParser()
		parser.RegisterFunction("foo", "bar")
		parser.RegisterFunction("foo", "baz")
		parser.RegisterFunction("foo", "+")
		parser.RegisterFunction("zoo", "-")

		Convey("It can return all possible symbols available", func() {
			So(
				parser.Symbols(), ShouldResemble,
				[]string{"foo.bar", "foo.baz", "foo.+", "zoo.-"},
			)
		})

		Convey("It can parse integer literals", func() {
			code := "1 98 -27"
			So(
				parser.Parse(code), ShouldResemble,
				InstructionSet{
					NewInstruction("integer", "1", ""),
					NewInstruction("integer", "98", ""),
					NewInstruction("integer", "-27", ""),
				},
			)
		})

		Convey("It can parse float literals", func() {
			code := "0.528 3.14 -0.189"
			So(
				parser.Parse(code), ShouldResemble,
				InstructionSet{
					NewInstruction("float", "0.528", ""),
					NewInstruction("float", "3.14", ""),
					NewInstruction("float", "-0.189", ""),
				},
			)
		})

		Convey("It can parse boolean literals", func() {
			code := "true false"
			So(
				parser.Parse(code), ShouldResemble,
				InstructionSet{
					NewInstruction("boolean", "true", ""),
					NewInstruction("boolean", "false", ""),
				},
			)
		})

		Convey("It can parse known data stack functions", func() {
			code := "foo.bar foo.baz foo.+"
			So(
				parser.Parse(code), ShouldResemble,
				InstructionSet{
					NewInstruction("foo", "foo.bar", "bar"),
					NewInstruction("foo", "foo.baz", "baz"),
					NewInstruction("foo", "foo.+", "+"),
				},
			)
		})

		Convey("It will ignore unknown data stacks and functions", func() {
			code := "foo.- fool.bar foo.bar"
			So(
				parser.Parse(code), ShouldResemble,
				InstructionSet{
					NewInstruction("foo", "foo.bar", "bar"),
				},
			)
		})
	})
}
