package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParser(t *testing.T) {
	Convey("Given a parser", t, func() {
		parser := NewParser()
		parser.Functions["foo"] = map[string]bool{"bar": true, "baz": true, "+": true}

		Convey("It can parse integer literals", func() {
			code := "1 98 -27"
			So(
				parser.Parse(code), ShouldResemble,
				[]Instruction{
					Instruction{"integer", "1", ""},
					Instruction{"integer", "98", ""},
					Instruction{"integer", "-27", ""},
				},
			)
		})

		Convey("It can parse float literals", func() {
			code := "0.528 3.14 -0.189"
			So(
				parser.Parse(code), ShouldResemble,
				[]Instruction{
					Instruction{"float", "0.528", ""},
					Instruction{"float", "3.14", ""},
					Instruction{"float", "-0.189", ""},
				},
			)
		})

		Convey("It can parse boolean literals", func() {
			code := "true false"
			So(
				parser.Parse(code), ShouldResemble,
				[]Instruction{
					Instruction{"boolean", "true", ""},
					Instruction{"boolean", "false", ""},
				},
			)
		})

		Convey("It can parse known data stack functions", func() {
			code := "foo.bar foo.baz foo.+"
			So(
				parser.Parse(code), ShouldResemble,
				[]Instruction{
					Instruction{"foo", "foo.bar", "bar"},
					Instruction{"foo", "foo.baz", "baz"},
					Instruction{"foo", "foo.+", "+"},
				},
			)
		})

		Convey("It will ignore unknown data stacks and functions", func() {
			code := "foo.- fool.barz foo.bar"
			So(
				parser.Parse(code), ShouldResemble,
				[]Instruction{
					Instruction{"foo", "foo.bar", "bar"},
				},
			)
		})
	})
}
