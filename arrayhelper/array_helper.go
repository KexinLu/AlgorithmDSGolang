package arrayhelper

import (
	"bytes"
	"fmt"
)

func PrintArray(arr []interface{}, ptrs []int) {
	var buffer bytes.Buffer
	var markerBuffer bytes.Buffer

	for idx, val := range arr {
		buffer.WriteString(fmt.Sprintf("%v,", val))
		if arrayContains(ptrs, idx) {
			markerBuffer.WriteByte('^')
			markerBuffer.WriteByte(' ')
		} else {
			markerBuffer.WriteByte(' ')
			markerBuffer.WriteByte(' ')
		}
	}

	fmt.Println(buffer.String())
	fmt.Println(markerBuffer.String())
}

func PrintIntArray(arr []int, ptrs ...int) {
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}

	PrintArray(s, ptrs)
}

func PrintByteArray(arr []byte, ptrs ...int) {
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}

	PrintArray(s, ptrs)
}

func IntArrayEquals(arr1, arr2 []int) bool {
	for idx, v := range arr1 {
		if arr2[idx] != v {
			return false
		}
	}

	return true
}

func arrayContains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func PrintStringArray(arr []string, ptrs ...int) {
	var buffer bytes.Buffer
	var markerBuffer bytes.Buffer

	for idx, val := range arr {
		buffer.WriteString(fmt.Sprintf("%v,", val))
		curLen := len(val)
		if arrayContains(ptrs, idx) {
			markerBuffer.WriteByte('^')
			for i := 0; i < curLen; i++ {
				markerBuffer.WriteByte(' ')
			}
		} else {
			for i := 0; i <= curLen; i++ {
				markerBuffer.WriteByte(' ')
			}
		}
	}

	fmt.Println(buffer.String())
	fmt.Println(markerBuffer.String())
}
