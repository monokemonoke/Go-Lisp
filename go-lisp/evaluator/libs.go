package evaluator

import (
	"fmt"
	"strconv"

	"github.com/monokemonoke/go-lisp/go-lisp/reader"
)

func add(args []any) any {
	fmt.Printf("debug\n")
	fmt.Printf("%#v\n", args)

	sum := 0
	for _, arg := range args {
		x, _ := strconv.Atoi(arg.(reader.Token).String())
		sum += x
	}
	return sum
}
