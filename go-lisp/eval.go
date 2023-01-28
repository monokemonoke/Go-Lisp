package golisp

import (
	"fmt"
	"strconv"
)

func Eval(ast Ast) any {
	astExpression := ast.(*astExpression)

	switch astExpression.operand {
	case "+":
		return add(astExpression.args)
	case "-":
		return minus(astExpression.args)
	default:
		return fmt.Sprintf("got unexpected operator %s\n", astExpression.operand)
	}
}

func add(args []string) int {
	sum := 0
	for _, val := range args {
		num, _ := strconv.Atoi(val)
		sum += num
	}
	return sum
}

func minus(args []string) int {
	x, _ := strconv.Atoi(args[0])
	for _, val := range args[1:] {
		num, _ := strconv.Atoi(val)
		x -= num
	}
	return x
}
