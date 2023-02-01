package main

import (
	"bufio"
	"fmt"
	"os"

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

		fmt.Printf("%#v\n", tokens)
	}
}
