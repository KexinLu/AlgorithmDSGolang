package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stk := NewStack()
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(5)

	stk.Print()
}
