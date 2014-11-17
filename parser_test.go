package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParser(t *testing.T) {
	Convey("Given a parser", t, func() {
		parser := NewParser()
		parser.Functions["integer"] = []string{"foo", "bar"}
		parser.Functions["boolean"] = []string{"+", "-"}
		parser.Functions["float"] = []string{"pizza", "pasta"}

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
	})
}
