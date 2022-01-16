package strings

import "testing"

func TestLoopThrough(t *testing.T) {
	input := "hello, this is a test, 你好"
	LoopThrough(input)
}

func TestStringBuilder(t *testing.T) {
	input := "hello, this is a test, 你好"
	StringBuilder(input)
}
