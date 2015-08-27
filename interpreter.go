// Package spogoto is an implementation of the stack-based language Spogoto.
package spogoto

import (
	"fmt"
	"math/rand"
	"time"
)

// Interpreter interprets Spogoto code.
type Interpreter interface {
	RandInt() int64
	RandFloat() float64
	Run(Code) RunSet
}

// Options sets the Instruction options changing its behavior depending
// on what values are set on the fields.
type Options struct {

	// MaxInstructions is the maximum number of total instruction executions
	MaxInstructions int64
}

// DefaultOptions is the default set of options.
var DefaultOptions = Options{
	MaxInstructions: 100,
}

type Rand interface {
	Int63n(int64) int64
	Float64() float64
}

type interpreter struct {
	Rand    Rand
	Parser  *Parser
	Options Options
}

// Run executes a Spogoto code string and returns a RunSet as result.
func (i *interpreter) Run(code Code) (r RunSet) {
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

// RandomCodeArray generates a random code of the specified length.
func (i *interpreter) RandomCode(length int64) Code {
	var code = Code{}
	var k int64
	for k = 0; k < length; k++ {
		code = append(code, i.RandomInstruction())
	}

	return code
}

// RandInt generates a random integer between 0 and 9.
func (i *interpreter) RandInt() int64 {
	return i.Rand.Int63n(10)
}

// RandFloat generates a random float number between 0 and 1.
func (i *interpreter) RandFloat() float64 {
	return i.Rand.Float64()
}

// RandomInstruction generates a random instruction. An Instruction can either be
// a literal or a function.
func (i *interpreter) RandomInstruction() string {
	if i.RandFloat() < 0.3 {
		return i.RandomLiteral([]string{"integer", "float", "boolean"}[i.Rand.Int63n(3)])
	}
	return i.RandomSymbol()
}

// RandomLiteral generates a random literal instruction of type t.
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

// RandomSymbol generates a random DataSet or Cursor function.
func (i *interpreter) RandomSymbol() string {
	symbols := i.Parser.Symbols()
	idx := i.Rand.Int63n(int64(len(symbols)))
	return symbols[idx]
}

func (i *interpreter) randomBoolean() string {
	if i.RandFloat() > 0.5 {
		return "true"
	}
	return "false"
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
		for fn := range stack.Functions() {
			p.RegisterFunction(t, fn)
		}
	}

	for fnc := range r.CursorCommands() {
		p.RegisterFunction("cursor", fnc)
	}

	i.Parser = p
}

// NewInterpreter constructs a new Intepreter configured with options.
func NewInterpreter(options Options) *interpreter {
	i := &interpreter{
		Rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Options: options,
	}
	i.setupParser(i.createRunSet())

	return i
}

func (i *interpreter) createRunSet() *runset {
	return NewRunSet(i)
}
