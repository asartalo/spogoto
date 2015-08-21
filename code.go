package spogoto

import (
	"regexp"
	"strings"
)

type Code []string

func (c Code) String() string {
	return strings.Join([]string(c), " ")
}

func CodeFromString(str string) Code {
	re := regexp.MustCompile("\\s+")
	return Code(re.Split(str, -1))
}
