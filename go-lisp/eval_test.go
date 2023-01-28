package golisp

import "testing"

func TestEval(t *testing.T) {
	testCases := []struct {
		desc   string
		src    Ast
		expect any
	}{
		{
			desc:   "basic add expression",
			src:    Parse(Read("(+ 1 1)")),
			expect: 2,
		},
		{
			desc:   "basic add expression with negative number",
			src:    Parse(Read("(+ -1 5)")),
			expect: 4,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Eval(tC.src)

			if tC.expect != actual {
				t.Errorf("want %+v but got %+v\n", tC.expect, actual)
			}
		})
	}
}
