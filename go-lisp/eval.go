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
