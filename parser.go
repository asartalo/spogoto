package spogoto

import (
	"fmt"
	"regexp"
)

type Instruction struct {
	Type     string
	Value    string
	Function string
}

type Parser struct {
	Functions map[string][]string
}

func (p *Parser) Parse(code string) []Instruction {
	re := regexp.MustCompile("\\s+")
	raw := re.Split(code, -1)
	i := []Instruction{}
	fmt.Println(raw)
	for _, item := range raw {
		i = append(i, p.ParseItem(item))
	}
	return i
}

func (p *Parser) ParseItem(item string) Instruction {
	var t string
	if item == "true" || item == "false" {
		t = "boolean"
	} else if regexp.MustCompile(`^\-?\d+$`).MatchString(item) {
		t = "integer"
	} else if regexp.MustCompile(`^\-?\d+\.\d+$`).MatchString(item) {
		t = "float"
	}

	return Instruction{t, item, ""}
}

func NewParser() *Parser {
	return &Parser{make(map[string][]string)}
}
