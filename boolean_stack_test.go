package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBooleanStackFunctions(t *testing.T) {

	testData := spogotoTestData{
		tbooleansOnly("=", []bool{true, false}, []bool{false}),
		tbooleansOnly("=", []bool{false, false}, []bool{true}),
		tbooleansOnly("and", []bool{true, false}, []bool{false}),
		tbooleansOnly("and", []bool{false, false}, []bool{false}),
		tbooleansOnly("and", []bool{true, true}, []bool{true}),
		tbooleansOnly("or", []bool{true, true}, []bool{true}),
		tbooleansOnly("or", []bool{false, true}, []bool{true}),
		tbooleansOnly("or", []bool{true, false}, []bool{true}),
		tbooleansOnly("or", []bool{false, false}, []bool{false}),
		tbooleansOnly("xor", []bool{true, true}, []bool{false}),
		tbooleansOnly("xor", []bool{false, true}, []bool{true}),
		tbooleansOnly("xor", []bool{true, false}, []bool{true}),
		tbooleansOnly("xor", []bool{false, false}, []bool{false}),
		tbooleansOnly("not", []bool{true}, []bool{false}),
		tbooleansOnly("not", []bool{false}, []bool{true}),
		{
			"frominteger",
			[]int64{1}, []int64{},
			[]bool{}, []bool{true},
			[]float64{}, []float64{},
		},
		{
			"frominteger",
			[]int64{0}, []int64{},
			[]bool{}, []bool{false},
			[]float64{}, []float64{},
		},
		{
			"fromfloat",
			[]int64{}, []int64{},
			[]bool{}, []bool{true},
			[]float64{1.2}, []float64{},
		},
		{
			"fromfloat",
			[]int64{}, []int64{},
			[]bool{}, []bool{false},
			[]float64{0.0}, []float64{},
		},

		// Empties
		tnothingHappens("="),
		tnothingHappens("and"),
		tnothingHappens("or"),
		tnothingHappens("not"),
		tnothingHappens("xor"),
		tnothingHappens("frominteger"),
		tnothingHappens("fromfloat"),
	}

	for _, d := range testData {
		i := NewInterpreter()
		r := NewRunSet(i)

		Convey(tPrimaryMessage("boolean", d), t, func() {
			integerStack := tGenericDataStack(int64Elements(d.intsBefore))
			floatStack := tGenericDataStack(float64Elements(d.floatsBefore))
			r.RegisterStack("integer", integerStack)
			r.RegisterStack("float", floatStack)

			s := NewBooleanStack(d.boolsBefore)
			Convey("It shouldn't panic", func() {
				So(func() { s.Call(d.fn, r, i) }, ShouldNotPanic)

				Convey(fmt.Sprintf("Boolean elements should be %v", d.boolsAfter), func() {
					So(s.elements, ShouldResemble, boolElements(d.boolsAfter))
				})

				Convey(tIntMessaging(d.intsBefore, d.intsAfter), func() {
					So(integerStack.elements, ShouldResemble, int64Elements(d.intsAfter))
				})
				Convey(tFloatMessaging(d.floatsBefore, d.floatsAfter), func() {
					So(floatStack.elements, ShouldResemble, float64Elements(d.floatsAfter))
				})

			})

		})

	}
}

func TestOtherBooleanStackFeatures(t *testing.T) {
	Convey("Given an empty boolean stack", t, func() {
		s := NewBooleanStack([]bool{})

		Convey("When pushing literals", func() {
			Convey("It should push proper boolean literal", func() {
				s.PushLiteral("true")
				So(s.Pop(), ShouldEqual, true)
			})

			Convey("It should not push improper boolean literal", func() {
				s.PushLiteral("rrr")
				So(s.Pop(), ShouldBeNil)
			})
		})

	})
}
