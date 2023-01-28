package main

import (
	"bufio"
	"fmt"
	"os"

	golisp "github.com/monokemonoke/go-lisp/go-lisp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		if ok := scanner.Scan(); !ok {
			break
		}
		program := scanner.Text()

		fmt.Println(
			golisp.Eval(
				golisp.Parse(
					golisp.Read(
						program,
					),
				),
			),
		)
	}
}
