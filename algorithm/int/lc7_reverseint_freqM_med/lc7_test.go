package lc7

import "testing"

func TestReverseInt(t *testing.T) {
	tcs := []struct {
		input int
		exp   int
	}{
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			120,
			21,
		},
		{
			1534236469,
			0,
		},
	}

	for _, tc := range tcs {
		rst := reverse(tc.input)
		if tc.exp != rst {
			t.Errorf("expecting %v but got %v", tc.exp, rst)
		}
	}

}
