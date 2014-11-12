package spogoto

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFloatStackFunctions(t *testing.T) {

	testData := spogotoTestData{
		{
			">",
			[]int64{}, []int64{},
			[]bool{}, []bool{true},
			[]float64{1.0, 6.0, 2.0}, []float64{1.0},
		},
		{
			">",
			[]int64{}, []int64{},
			[]bool{}, []bool{false},
			[]float64{1.0, 2.0, 6.0}, []float64{1.0},
		},
		{
			"<",
			[]int64{}, []int64{},
			[]bool{}, []bool{false},
			[]float64{1.0, 6.0, 2.0}, []float64{1.0},
		},
		{
			"<",
			[]int64{}, []int64{},
			[]bool{}, []bool{true},
			[]float64{1.0, 2.0, 6.0}, []float64{1.0},
		},
	}

	for _, d := range testData {
		i := NewInterpreter()
		boolStack := NewDataStack(boolElements(d.boolsBefore), FunctionMap{})
		integerStack := NewDataStack(int64Elements(d.intsBefore), FunctionMap{})
		i.RegisterStack("boolean", boolStack)
		i.RegisterStack("integer", integerStack)

		Convey(fmt.Sprintf("Call()ing '%s' with float stack %v", d.fn, d.floatsBefore), t, func() {
			s := NewFloatStack(d.floatsBefore)
			s.Call(d.fn, i)
			So(s.elements, ShouldResemble, float64Elements(d.floatsAfter))
			So(boolStack.elements, ShouldResemble, boolElements(d.boolsAfter))
			So(integerStack.elements, ShouldResemble, int64Elements(d.intsAfter))
		})

	}
}
