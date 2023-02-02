package main

import (
	"bufio"
	"fmt"
	"os"

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
		fmt.Printf("tokens\n")
		fmt.Printf("%#v\n", tokens)

		p := parser.NewParser(tokens)
		asts := p.Parse()
		fmt.Printf("asts\n")
		fmt.Printf("%#v\n", asts)
	}
}
