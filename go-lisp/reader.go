package golisp

import "strings"

func Read(program string) []string {
	return strings.Split(program, " ")
}
