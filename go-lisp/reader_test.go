package golisp

import "testing"

func TestRead(t *testing.T) {
	testCases := []struct {
		desc   string
		src    string
		expect []string
	}{
		{
			desc:   "basic test case 1",
			src:    "Hello",
			expect: []string{"Hello"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Read(tC.src)
			if len(tC.expect) != len(actual) {
				t.Errorf("want %d items but got %d items", len(tC.expect), len(actual))
				return
			}

			for i := range actual {
				if tC.expect[i] != actual[i] {
					t.Errorf("want %+v but got %+v\n", tC.expect, actual)
				}
			}
		})
	}
}
