package lc8

import "testing"

func TestMyAtoi(t *testing.T) {
	tcs := []struct {
		in  string
		exp int
	}{

		{
			"42",
			42,
		},
		{
			"   -42",
			-42,
		},
		{
			"4193 with words",
			4193,
		},
		{
			"4193with words",
			4193,
		},
		{
			"words and 987",
			0,
		},
		{
			"-91283472332",
			-2147483648,
		},
		{
			"+1",
			1,
		},
		{
			"+-12",
			0,
		},
		{
			"21474836460",
			21474836467,
		},
		{
			"0000000-42",
			0,
		},
		{
			"   +0 123",
			0,
		},
	}

	for _, tc := range tcs {
		rst := myAtoi(tc.in)
		if tc.exp != rst {
			t.Errorf("expecting %v but got %v", tc.exp, rst)
		}
	}

}
