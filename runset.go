package spogoto

// CursorCommands are functions that operate on the Cursor manipulating
// its position.
type CursorCommands map[string]func(RunSet)

// RunSet is a container for a Spogoto code's execution environment.
// The RunSet contains the DataStacks that the code will operate on as
// well as other information of regarding a code's execution.
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

// Cursor is a representation of a pointer pointing to the current
// Instruction on which the interpreter will have to execute.
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

// NewRunSet creates a RunSet.
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

// Cursor returns the Cursor.
func (r *runset) Cursor() *Cursor {
	return &r.cursor
}

// IncrementInstructionCount increments the InstructionCount of course.
func (r *runset) IncrementInstructionCount() {
	r.instructionCount++
}

// InstructionCount returns the total number of Instructions executed for
// a single run.
func (r *runset) InstructionCount() int64 {
	return r.instructionCount
}

// CursorCommand executes cursor-related functions.
func (r *runset) CursorCommand(fn string) {
	theFunc, ok := r.cursorCommands[fn]
	if ok {
		theFunc(r)
	}
}

// CursorCommands returns all CursorCommands available.
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

// DataStacks returns all the DataStacks registered for this RunSet.
func (r *runset) DataStacks() map[string]DataStack {
	return r.dataStacks
}

func instructionCount(r RunSet) int64 {
	return int64(len(r.Cursor().Instructions))
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
		r.Cursor().Position = instructionCount(r)
	}

	commands["endif"] = func(r RunSet) {
		if r.Bad("boolean", 1) {
			return
		}
		if r.Stack("boolean").Pop().(bool) {
			commands["end"](r)
		}
	}

	commands["goto"] = func(r RunSet) {
		if r.Bad("integer", 1) {
			return
		}
		pos := r.Stack("integer").Pop().(int64)
		if pos < 0 || pos > instructionCount(r) {
			return
		}
		r.Cursor().Position = int64(pos - 1)
	}

	commands["gotoif"] = func(r RunSet) {
		if r.Ok("boolean", 1) && r.Stack("boolean").Pop().(bool) {
			commands["goto"](r)
		}
	}

	rs.cursorCommands = commands
}
