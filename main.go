package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/monokemonoke/go-lisp/go-lisp/evaluator"
	"github.com/monokemonoke/go-lisp/go-lisp/parser"
	"github.com/monokemonoke/go-lisp/go-lisp/reader"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		if ok := scanner.Scan(); !ok {
			break
		}
		program := scanner.Text()

		tokens := reader.NewReader(program).Read()

		p := parser.NewParser(tokens)
		asts := p.Parse()

		res := evaluator.Eval(asts)
		fmt.Printf("%#v\n", res)
	}
}
