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

type nodeAtom struct {
	value any
}

type nodeExp struct {
	operand string
	args    []Node
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
		return nodeAtom{value: top}
	}

	var ast nodeExp
	opToken, _ := p.seek()
	ast.operand = string(opToken)

	for {
		token, _ := p.seek()
		if token == ")" {
			return ast
		}

		switch token {
		case "(":
			p.position--
			ast.args = append(ast.args, p.Parse())
		default:
			ast.args = append(ast.args, nodeAtom{value: token})
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
