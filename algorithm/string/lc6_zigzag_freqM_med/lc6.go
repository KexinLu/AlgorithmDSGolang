package lc6

import "bytes"

/**
https://leetcode.com/problems/zigzag-conversion/
**/

// stor by row
func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	bufs := make([]*bytes.Buffer, numRows)
	for i := 0; i < numRows; i++ {
		bufs[i] = new(bytes.Buffer)
	}
	bufIdx := 0
	for i := 0; i < len(s); i++ {
		bufs[bufIdx].WriteByte(s[i])
		if (i/(numRows-1))&1 == 1 {
			bufIdx--
		} else {
			bufIdx++
		}
	}

	rst := ""
	for _, v := range bufs {
		rst += v.String()
	}
	return rst

}
