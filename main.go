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
		scanner.Scan()
		program := scanner.Text()

		fmt.Println(program)
	}
}
