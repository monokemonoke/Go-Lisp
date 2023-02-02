package evaluator

import (
	"github.com/monokemonoke/go-lisp/go-lisp/parser"
)

func Eval(asts parser.Node) any {
	switch v := asts.(type) {
	case parser.NodeAtom:
		return v.Value
	case parser.NodeExp:
		fn := findFuncs(v.Operand)
		var evalArgs []any
		for _, arg := range v.Args {
			evalArgs = append(evalArgs, Eval(arg))
		}
		return fn(evalArgs)
	default:
		return nil
	}
}

type Func func(args []any) any

func NewFuncMap() map[string]Func {
	m := make(map[string]Func)
	m["+"] = add
	m["-"] = minus
	m["*"] = multi
	m["/"] = div
	return m
}

func findFuncs(id string) Func {
	FuncMap := NewFuncMap()
	fn, ok := FuncMap[id]
	if ok {
		return fn
	}

	return func(args []any) any {
		return nil
	}
}
