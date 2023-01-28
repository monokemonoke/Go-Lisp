package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		if ok := scanner.Scan(); !ok {
			break
		}
		program := scanner.Text()

		fmt.Println(program)
	}
}
