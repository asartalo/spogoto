package spogoto

import (
	"regexp"
	"strings"
)

// Instruction is a unit of code signifying the type it operates on,
// the value literal of the instruction, the function to call if present,
// and the number of runs or executions that the instruction has been called.
type Instruction struct {
	Type     string
	Value    string
	Function string
	Runs     int
}

// NewInstruction creates a new Instruction.
func NewInstruction(t string, val string, fn string) Instruction {
	return Instruction{t, val, fn, 0}
}

// InstructionSet is a list of Instructions.
type InstructionSet []Instruction

// Parser parses string codes into an InstructionSet.
type Parser struct {
	Functions map[string]map[string]bool
	symbols   []string
}

// FunctionRegistered returns true if a function of a type has been registered.
func (p *Parser) FunctionRegistered(t string, fn string) bool {
	funcs, ok := p.Functions[t]
	if ok {
		return funcs[fn]
	}
	return false
}

// Parse parses string into an InstructionSet.
func (p *Parser) Parse(code string) InstructionSet {
	re := regexp.MustCompile("\\s+")
	raw := re.Split(code, -1)
	i := InstructionSet{}
	for _, item := range raw {
		parsed := p.ParseItem(item)
		if parsed.Type != "" {
			i = append(i, parsed)
		}
	}
	return i
}

// ParseItem parses a single string instruction into an Instruciton.
func (p *Parser) ParseItem(item string) Instruction {
	var t string
	var fn string
	if item == "true" || item == "false" {
		t = "boolean"
	} else if regexp.MustCompile(`^\-?\d+$`).MatchString(item) {
		t = "integer"
	} else if regexp.MustCompile(`^\-?\d+\.\d+$`).MatchString(item) {
		t = "float"
	} else if regexp.MustCompile(`^\w+\.[^\.]+$`).MatchString(item) {
		s := strings.Split(item, ".")
		t = s[0]
		fn = s[1]
		if !p.FunctionRegistered(t, fn) {
			t = ""
			fn = ""
		}

	}

	return NewInstruction(t, item, fn)
}

// RegisterFunction registers a function of a type t and adds it to its
// list of available functions.
func (p *Parser) RegisterFunction(t string, fn string) {
	m, ok := p.Functions[t]
	if !ok {
		p.Functions[t] = map[string]bool{}
		m = p.Functions[t]
	}
	m[fn] = true
	p.symbols = append(p.symbols, t+"."+fn)

}

// Symbols return all available symbols/functions.
func (p *Parser) Symbols() []string {
	return p.symbols
}

// NewParser creates a new Parser.
func NewParser() *Parser {
	return &Parser{make(map[string]map[string]bool), []string{}}
}
