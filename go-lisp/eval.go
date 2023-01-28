package golisp

import (
	"fmt"
	"strconv"
)

func Eval(ast Ast) any {
	astExpression := ast.(*astExpression)

	switch astExpression.operand {
	case "+":
		sum := 0
		for _, val := range astExpression.args {
			num, _ := strconv.Atoi(val)
			sum += num
		}
		return sum
	default:
		return fmt.Sprintf("got unexpected operator %s\n", astExpression.operand)
	}
}
