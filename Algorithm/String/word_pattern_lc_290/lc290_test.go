package lc290

import "testing"

func TestWordPattern(t *testing.T) {
	tcs := []struct {
		pattern string
		input   string
		exp     bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"abc", "dog cat dog", false},
	}

	for idx, tc := range tcs {
		rst := wordPattern(tc.pattern, tc.input)
		if rst != tc.exp {
			t.Errorf("expected %v is different from the result %v for testcase %d", tc.exp, rst, idx)
		}
	}
}
