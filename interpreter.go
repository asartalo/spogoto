package spogoto

import (
	"math/rand"
)

// Interpreter interprets Spogoto code.
type Interpreter interface {
	RegisterStack(string, DataStack)
	Stack(string) DataStack
	Ok(string, int64) bool
	Bad(string, int64) bool
	RandInt() int64
	RandFloat() float64
}

type Cursor struct {
	Position int64
}

type interpreter struct {
	DataStacks     map[string]DataStack
	Rand           *rand.Rand
	Parser         *Parser
	Cursor         Cursor
	CursorCommands map[string]func(*interpreter)
}

func (i *interpreter) Run(code string) {
	instructions := i.Parser.Parse(code)
	inCount := int64(len(instructions))
	for i.Cursor.Position < inCount {
		instruction := instructions[i.Cursor.Position]
		t := instruction.Type
		fn := instruction.Function
		if t == "" {
			// No type so noop
			i.Cursor.Position++
			continue
		}

		if fn == "" {
			// Literal type
			i.Stack(t).PushLiteral(instruction.Value)
		} else if t == "cursor" {
			i.CursorCommand(fn)
		} else {
			// It's calling a function
			i.Stack(t).Call(fn, i)
		}

		i.Cursor.Position++
	}
}

func (i *interpreter) CursorCommand(fn string) {
	theFunc, ok := i.CursorCommands[fn]
	if ok {
		theFunc(i)
	}
}

// RegisterStack adds a stack to the available DataStacks identified by name.
func (i *interpreter) RegisterStack(name string, stack DataStack) {
	i.DataStacks[name] = stack
}

// Stack returns the stack registered with name.
func (i *interpreter) Stack(name string) DataStack {
	s, ok := i.DataStacks[name]
	if !ok {
		s = &NullDataStack{}
	}

	return s
}

// Ok returns true if stack is available and has the number of elements
// required.
func (i *interpreter) Ok(name string, count int64) bool {
	return !i.Bad(name, count)
}

// Bad returns false if stack is available and has the number of elements
// required.
func (i *interpreter) Bad(name string, count int64) bool {
	return i.Stack(name).Lack(count)
}

// RandInt generates a random integer between 0 and 9
func (i *interpreter) RandInt() int64 {
	return i.Rand.Int63n(10)
}

// RandFloat generates a random float number between 0 and 1
func (i *interpreter) RandFloat() float64 {
	return i.Rand.Float64()
}

func (i *interpreter) setupParser(p *Parser) {
	for t, stack := range i.DataStacks {
		for fn, _ := range stack.Functions() {
			p.RegisterFunction(t, fn)
		}
	}

	i.Parser = p
}

// NewInterpreter constructs a new Intepreter.
func NewInterpreter() *interpreter {
	p := NewParser()
	i := &interpreter{
		DataStacks: map[string]DataStack{
			"integer": NewIntegerStack([]int64{}),
			"float":   NewFloatStack([]float64{}),
			"boolean": NewBooleanStack([]bool{}),
		},
		Rand: rand.New(rand.NewSource(rand.Int63())),
	}
	addCursorCommands(i)
	i.setupParser(p)

	return i
}

func addCursorCommands(in *interpreter) {
	commands := make(map[string]func(*interpreter))

	commands["skipif"] = func(i *interpreter) {
		if i.Bad("boolean", 1) {
			return
		}
		if i.Stack("boolean").Pop().(bool) {
			i.Cursor.Position++
		}
	}

	in.CursorCommands = commands
}
