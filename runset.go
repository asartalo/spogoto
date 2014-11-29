package spogoto

type CursorCommands map[string]func(RunSet)

type RunSet interface {
	RegisterStack(string, DataStack)
	Stack(string) DataStack
	Ok(string, int64) bool
	Bad(string, int64) bool
	DataStacks() map[string]DataStack
	Cursor() *Cursor
	CursorCommand(string)
	CursorCommands() CursorCommands
	IncrementInstructionCount()
	InstructionCount() int64
}

type Cursor struct {
	Position     int64
	Instructions InstructionSet
}

type runset struct {
	dataStacks       map[string]DataStack
	cursor           Cursor
	cursorCommands   map[string]func(RunSet)
	instructionCount int64
}

func NewRunSet(i Interpreter) *runset {
	r := &runset{
		dataStacks: map[string]DataStack{
			"integer": NewIntegerStack([]int64{}),
			"float":   NewFloatStack([]float64{}),
			"boolean": NewBooleanStack([]bool{}),
		},
	}
	addCursorCommands(r)

	return r
}

func (r *runset) Cursor() *Cursor {
	return &r.cursor
}

func (r *runset) IncrementInstructionCount() {
	r.instructionCount++
}

func (r *runset) InstructionCount() int64 {
	return r.instructionCount
}

func (r *runset) CursorCommand(fn string) {
	theFunc, ok := r.cursorCommands[fn]
	if ok {
		theFunc(r)
	}
}

func (r *runset) CursorCommands() CursorCommands {
	return r.cursorCommands
}

// RegisterStack adds a stack to the available DataStacks identified by name.
func (r *runset) RegisterStack(name string, stack DataStack) {
	r.dataStacks[name] = stack
}

// Stack returns the stack registered with name.
func (r *runset) Stack(name string) DataStack {
	s, ok := r.dataStacks[name]
	if !ok {
		s = &NullDataStack{}
	}

	return s
}

// Ok returns true if stack is available and has the number of elements
// required.
func (r *runset) Ok(name string, count int64) bool {
	return !r.Bad(name, count)
}

// Bad returns false if stack is available and has the number of elements
// required.
func (r *runset) Bad(name string, count int64) bool {
	return r.Stack(name).Lack(count)
}

func (r *runset) DataStacks() map[string]DataStack {
	return r.dataStacks
}

func addCursorCommands(rs *runset) {
	commands := make(CursorCommands)

	commands["skipif"] = func(r RunSet) {
		if r.Bad("boolean", 1) {
			return
		}
		if r.Stack("boolean").Pop().(bool) {
			r.Cursor().Position++
		}
	}

	commands["end"] = func(r RunSet) {
		r.Cursor().Position = int64(len(r.Cursor().Instructions))
	}

	commands["endif"] = func(r RunSet) {
		if r.Bad("boolean", 1) {
			return
		}
		if r.Stack("boolean").Pop().(bool) {
			r.Cursor().Position = int64(len(r.Cursor().Instructions))
		}
	}

	rs.cursorCommands = commands
}
