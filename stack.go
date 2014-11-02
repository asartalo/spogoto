package spogoto

// Element represents elements added to a stack.
type Element interface{}

// Elements is a slice of Element(s).
type Elements []Element

// Stack is a stack of Elements. The top most element of the stack is at index 0.
type Stack interface {
	Size() int64
	IsEmpty() bool
	Has(n int64) bool
	Peek() Element
	Pop() Element
	Push(i Element)
	Swap()
	Flush()
	Rotate()
	Yank(idx int64)
	YankDup(idx int64)
	Shove(e Element, idx int64)
}

type stack struct {
	elements Elements
}

// NewStack creates a new stack filled with Elements.
func NewStack(elements Elements) stack {
	return stack{elements}
}

// Size returns the number of elements in the stack.
func (s *stack) Size() int64 {
	return int64(len(s.elements))
}

// IsEmpty returns true if there are no elements on the stack.
func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

// Has returns true if it has n or more elements in the stack.
func (s *stack) Has(n int64) bool {
	return n <= s.Size()
}

// Peek returns the topmost value on the stack or nil if there are no values
// on the stack.
func (s *stack) Peek() Element {
	if s.IsEmpty() {
		return nil
	}

	return s.elements[len(s.elements)-1]
}

// Pop returns the topmost value on the stack, removing it from the stack
// or nil if there are no values on the stack.
func (s *stack) Pop() Element {
	if s.IsEmpty() {
		return nil
	}

	i := s.Size() - 1
	e := s.elements[i]
	s.elements = s.elements[:i]
	return e
}

// Push adds the element to the top of the stack.
func (s *stack) Push(i Element) {
	s.elements = append(s.elements, i)
}

// Swap swaps the positions of the top two elements of the stack.
func (s *stack) Swap() {
	if s.Size() < 2 {
		return
	}

	last1 := s.Size() - 1
	last2 := last1 - 1
	s.elements[last1], s.elements[last2] = s.elements[last2], s.elements[last1]
}

// Flush empties the stack.
func (s *stack) Flush() {
	s.elements = Elements{}
}

// Rotate pulls the third element off the stack and places it on top.
func (s *stack) Rotate() {
	if s.Size() < 3 {
		return
	}

	l1 := s.Size() - 1
	l2 := l1 - 1
	l3 := l2 - 1
	s.elements[l1], s.elements[l2], s.elements[l3] =
		s.elements[l3], s.elements[l1], s.elements[l2]
}

func (s *stack) index(idx int64) int64 {
	return s.Size() - idx - 1
}

// Yank pulls an item of the specified index off the stack and places it on top.
func (s *stack) Yank(idx int64) {
	i := s.index(idx)
	// The s.Size() - 2 will be the index of the top most element so nothing to do
	if i > s.Size()-2 || i < 0 {
		return
	}

	e := s.elements[i]
	s.elements = append(
		append(s.elements[:i], s.elements[i+1:]...), e,
	)
}

// YankDup copies an item of the specified index and places the copy on top of the stack.
func (s *stack) YankDup(idx int64) {
	i := s.index(idx)
	if i > s.Size()-1 || i < 0 {
		return
	}

	e := s.elements[i]
	s.elements = append(s.elements, e)
}

// Shove inserts the item at the specified index.
func (s *stack) Shove(e Element, idx int64) {
	i := s.index(idx)
	if i > s.Size()-1 || i < 0 {
		return
	}

	s.elements = append(append(s.elements[:i], e), s.elements[i:]...)
}
