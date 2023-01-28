package golisp

import (
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		desc   string
		src    []string
		expect Ast
	}{
		{
			desc:   "basic add",
			src:    []string{"(", "+", "1", "1", ")"},
			expect: &astExpression{operand: "+", args: []string{"1", "1"}},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Parse(tC.src)

			if tC.expect.NodeType() != actual.NodeType() {
				t.Errorf("want %s but got %s\n", tC.expect.NodeType(), actual.NodeType())
			}

			actualExpression, ok := actual.(*astExpression)
			if !ok {
				t.Errorf("actual is not astExpression")
			}

			if tC.expect.(*astExpression).operand != actualExpression.operand {
				t.Errorf("want %s but got %s\n",
					tC.expect.(*astExpression).operand, actualExpression.operand)
			}
		})
	}
}
