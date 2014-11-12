package spogoto

type spogotoTestData []struct {
	fn           string
	intsBefore   []int64
	intsAfter    []int64
	boolsBefore  []bool
	boolsAfter   []bool
	floatsBefore []float64
	floatsAfter  []float64
}

func int64Elements(ints []int64) Elements {
	elements := Elements{}
	for _, v := range ints {
		elements = append(elements, int64(v))
	}
	return elements
}

func float64Elements(ints []float64) Elements {
	elements := Elements{}
	for _, v := range ints {
		elements = append(elements, float64(v))
	}
	return elements
}

func boolElements(bools []bool) Elements {
	elements := Elements{}
	for _, v := range bools {
		elements = append(elements, bool(v))
	}
	return elements
}
