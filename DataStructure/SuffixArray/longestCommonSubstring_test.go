package suffixArray

import (
	"fmt"
	"testing"
)

func TestCreateSuffixArray(t *testing.T) {
	rst := CreateSuffixArray("abcbca")
	expected := []string{
		"a",
		"abcbca",
		"bca",
		"bcbca",
		"ca",
		"cbca",
	}

	exp := fmt.Sprint(expected)
	rsts := fmt.Sprint(rst)
	if rsts != exp {
		t.Errorf("the suffix array result %v is different from expected %v", rsts, exp)
	}
}

func TestCreateLCPArray(t *testing.T) {
	suff := CreateSuffixArray("abcbca")
	lcp := CreateLCPArray(suff)
	expected := fmt.Sprint([]int{0, 1, 0, 2, 0, 1})
	rst := fmt.Sprint(lcp)

	if rst != expected {
		t.Errorf("the lcp array result %v is different from expected %v", rst, expected)
	}
}

func TestLongestRepeatedSubstring(t *testing.T) {
	rst := FindLongestRepeatedSubString("abccbabdabccdbabd")
	fmt.Println(rst)

}
