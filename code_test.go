package spogoto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCode(t *testing.T) {
	Convey("Given a string of code", t, func() {
		str := "true 4 cursor.gotoif 2 1 8"
		So(
			CodeFromString(str), ShouldResemble,
			Code{"true", "4", "cursor.gotoif", "2", "1", "8"},
		)

	})

	Convey("Given a code", t, func() {
		code := Code{"true", "4", "cursor.gotoif", "2", "1", "8"}
		So(
			code.String(), ShouldEqual,
			"true 4 cursor.gotoif 2 1 8",
		)
	})
}
