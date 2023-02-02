package parser

import (
	"errors"

	"github.com/monokemonoke/go-lisp/go-lisp/reader"
)

type Parser struct {
	position int
	tokens   []reader.Token
}

type Node interface{}

type NodeAtom struct {
	Value any
}

type NodeExp struct {
	Operand string
	Args    []Node
}

func NewParser(tokens []reader.Token) *Parser {
	return &Parser{
		position: 0,
		tokens:   tokens,
	}
}

func (p *Parser) Parse() any {
	top, _ := p.seek()
	if top != reader.Token("(") {
		return NodeAtom{Value: top}
	}

	var ast NodeExp
	opToken, _ := p.seek()
	ast.Operand = string(opToken)

	for {
		token, _ := p.seek()
		if token == ")" {
			return ast
		}

		switch token {
		case "(":
			p.position--
			ast.Args = append(ast.Args, p.Parse())
		default:
			ast.Args = append(ast.Args, NodeAtom{Value: any(token)})
		}
	}
}

func (p *Parser) seek() (reader.Token, error) {
	if p.position > len(p.tokens)-1 {
		return "", errors.New("end of tokens")
	}

	p.position++
	return p.tokens[p.position-1], nil
}
