package spogoto

import (
	"fmt"
	"math/rand"
)

// Interpreter interprets Spogoto code.
type Interpreter interface {
	RandInt() int64
	RandFloat() float64
	Run(string) RunSet
}

type Options struct {
	MaxInstructions int64
}

var DefaultOptions = Options{
	MaxInstructions: 100,
}

type interpreter struct {
	Rand    *rand.Rand
	Parser  *Parser
	Options Options
}

func (i *interpreter) Run(code string) (r RunSet) {
	r = i.createRunSet()
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
		if r.InstructionCount() > i.Options.MaxInstructions {
			break
		}
	}

	return r
}

// RandInt generates a random integer between 0 and 9
func (i *interpreter) RandInt() int64 {
	return i.Rand.Int63n(10)
}

// RandFloat generates a random float number between 0 and 1
func (i *interpreter) RandFloat() float64 {
	return i.Rand.Float64()
}

// RandomLiteral
func (i *interpreter) RandomLiteral(t string) string {
	switch t {
	case "integer":
		return i.randomInteger()
	case "float":
		return i.randomFloat()
	default:
		return i.randomBoolean()
	}
}

func (i *interpreter) RandomSymbol() string {
	symbols := i.Parser.Symbols()
	idx := i.Rand.Int63n(int64(len(symbols)))
	return symbols[idx]
}

func (i *interpreter) randomBoolean() string {
	if i.RandFloat() > 0.5 {
		return "true"
	} else {
		return "false"
	}
}

func (i *interpreter) randomInteger() string {
	var sign = ""
	if i.RandFloat() > 0.5 {
		sign = "-"
	}
	return fmt.Sprintf("%s%d", sign, i.RandInt())
}

func (i *interpreter) randomFloat() string {
	var sign = ""
	if i.RandFloat() > 0.5 {
		sign = "-"
	}
	return fmt.Sprintf("%s%f", sign, i.RandFloat())
}

func (i *interpreter) setupParser(r RunSet) {
	p := NewParser()
	for t, stack := range r.DataStacks() {
		for fn, _ := range stack.Functions() {
			p.RegisterFunction(t, fn)
		}
	}

	for fnc, _ := range r.CursorCommands() {
		p.RegisterFunction("cursor", fnc)
	}

	i.Parser = p
}

// NewInterpreter constructs a new Intepreter.
func NewInterpreter(options Options) *interpreter {
	i := &interpreter{
		Rand:    rand.New(rand.NewSource(rand.Int63())),
		Options: options,
	}
	i.setupParser(i.createRunSet())

	return i
}

func (i *interpreter) createRunSet() *runset {
	return NewRunSet(i)
}
