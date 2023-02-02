package evaluator

import (
	"strconv"

	"github.com/monokemonoke/go-lisp/go-lisp/reader"
)

func add(args []any) any {
	sum := 0
	for _, arg := range args {
		x, _ := strconv.Atoi(arg.(reader.Token).String())
		sum += x
	}
	return sum
}

func minus(args []any) any {
	x, _ := strconv.Atoi(args[0].(reader.Token).String())

	for _, arg := range args[1:] {
		y, _ := strconv.Atoi(arg.(reader.Token).String())
		x -= y
	}
	return x
}
