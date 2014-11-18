package spogoto

import (
	"regexp"
	"strings"
)

type Instruction struct {
	Type     string
	Value    string
	Function string
}

type Parser struct {
	Functions map[string]map[string]bool
}

func (p *Parser) FunctionRegistered(t string, fn string) bool {
	funcs, ok := p.Functions[t]
	if ok {
		return funcs[fn]
	}
	return false
}

func (p *Parser) Parse(code string) []Instruction {
	re := regexp.MustCompile("\\s+")
	raw := re.Split(code, -1)
	i := []Instruction{}
	for _, item := range raw {
		parsed := p.ParseItem(item)
		if parsed.Type != "" {
			i = append(i, parsed)
		}
	}
	return i
}

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

	return Instruction{t, item, fn}
}

func NewParser() *Parser {
	return &Parser{make(map[string]map[string]bool)}
}
