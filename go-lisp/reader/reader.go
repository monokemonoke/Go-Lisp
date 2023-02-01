package reader

import "strings"

type Token string

type Reader interface {
	Read() []Token
}

type reader struct {
	program string
}

func NewReader(program string) Reader {
	return reader{program: program}
}

func (r reader) Read() []Token {
	r.program = strings.ReplaceAll(r.program, "(", " ( ")
	r.program = strings.ReplaceAll(r.program, ")", " ) ")
	tmp := strings.Split(r.program, " ")

	var tokens []Token
	for _, val := range tmp {
		if val == "" {
			continue
		}
		tokens = append(tokens, Token(val))
	}
	return tokens
}
