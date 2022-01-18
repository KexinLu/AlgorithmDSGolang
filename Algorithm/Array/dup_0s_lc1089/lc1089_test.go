package lc1089

import (
	"fmt"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	input := []int{1, 0, 2, 3, 0, 4, 5, 0}
	exp := []int{1, 0, 0, 2, 3, 0, 0, 4}
	duplicateZeros(input)

	if fmt.Sprint(input) != fmt.Sprint(exp) {
		t.Errorf("expecting %v but got %v", exp, input)
	}
}
