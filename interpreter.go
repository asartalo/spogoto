package spogoto

import (
	"math/rand"
)

// Interpreter interprets Spogoto code.
type Interpreter interface {
	RandInt() int64
	RandFloat() float64
	Run(string) RunSet
}

type RunSet interface {
	RegisterStack(string, DataStack)
	Stack(string) DataStack
	Ok(string, int64) bool
	Bad(string, int64) bool
	DataStacks() map[string]DataStack
	Cursor() *Cursor
	CursorCommand(string)
	IncrementInstructionCount()
	InstructionCount() int64
}

type Cursor struct {
	Position     int64
	Instructions InstructionSet
}

type interpreter struct {
	Rand   *rand.Rand
	Parser *Parser
}

type runset struct {
	dataStacks       map[string]DataStack
	cursor           Cursor
	cursorCommands   map[string]func(RunSet)
	instructionCount int64
}

func (i *interpreter) Run(code string) (r RunSet) {
	r = i.createRunSet()
	if i.Parser == nil {
		i.setupParser(r)
	}
	instructions := i.Parser.Parse(code)
	inCount := int64(len(instructions))
	r.Cursor().Instructions = instructions
	for r.Cursor().Position < inCount {
		instruction := instructions[r.Cursor().Position]
		t := instruction.Type
		fn := instruction.Function
		if t == "" {
			// No type so noop
			r.Cursor().Position++
			continue
		}

		if fn == "" {
			// Literal type
			r.Stack(t).PushLiteral(instruction.Value)
		} else if t == "cursor" {
			r.CursorCommand(fn)
		} else {
			// It's calling a function
			r.Stack(t).Call(fn, r, i)
		}

		r.Cursor().Position++
		r.IncrementInstructionCount()
	}

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

// RandInt generates a random integer between 0 and 9
func (i *interpreter) RandInt() int64 {
	return i.Rand.Int63n(10)
}

// RandFloat generates a random float number between 0 and 1
func (i *interpreter) RandFloat() float64 {
	return i.Rand.Float64()
}

func (i *interpreter) setupParser(r RunSet) {
	p := NewParser()
	for t, stack := range r.DataStacks() {
		for fn, _ := range stack.Functions() {
			p.RegisterFunction(t, fn)
		}
	}

	i.Parser = p
}

// NewInterpreter constructs a new Intepreter.
func NewInterpreter() *interpreter {
	i := &interpreter{
		Rand: rand.New(rand.NewSource(rand.Int63())),
	}

	return i
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

func (i *interpreter) createRunSet() *runset {
	return NewRunSet(i)
}

func addCursorCommands(rs *runset) {
	commands := make(map[string]func(RunSet))

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
