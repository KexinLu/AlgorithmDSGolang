package arrayhelper

import "testing"

func TestPrintArray(t *testing.T) {
	PrintIntArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	PrintIntArray([]int{1, 2, 3, 4, 5, 6, 7}, 3, 5)
}

func TestPrintStringArray(t *testing.T) {
	PrintStringArray([]string{"my", "tests", "", "is", "not", "good"}, 1)
	PrintStringArray([]string{"my", "tests", "", "is", "not", "good"}, 1, 3)
}
