package golisp

type Ast interface {
	NodeType() string
}

type astExpression struct {
	operand string
	args    []string
}

func (a *astExpression) NodeType() string {
	return "astExpression"
}

func Parse(tokens []string) Ast {
	// 最初のカッコは読み飛ばす
	tokens = tokens[1:]

	var ast astExpression
	ast.operand = tokens[0]
	ast.args = tokens[1 : len(tokens)-1]
	return &ast
}
