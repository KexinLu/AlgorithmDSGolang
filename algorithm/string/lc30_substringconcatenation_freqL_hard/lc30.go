package lc30

import (
	"bytes"
)

// hashMap
func findSubstring(s string, words []string) []int {
	rst := []int{}
	if s == "" {
		return rst
	}
	if len(words) == 0 {
		return rst
	}
	wl := len(words[0])
	wc := len(words)
	wp := make(map[string]int)
	for _, word := range words {
		if _, ok := wp[word]; ok {
			wp[word]++
		} else {
			wp[word] = 1
		}
	}

	for i := 0; i < len(s); i++ {
		rst = append(rst, find(s, wp, wc, wl, i)...)
	}

	return rst
}

func find(s string, wp map[string]int, wc int, wl int, idx int) []int {
	if len(s[idx:]) < wl*wc {
		return []int{}
	}

	mp := make(map[string]int)
	for k, v := range wp {
		mp[k] = v
	}

	var buffer bytes.Buffer
	counter := 0
	for i := idx; i < len(s) && wc > 0; i++ {
		buffer.WriteByte(s[i])
		counter++
		if counter == wl {
			if val, ok := mp[buffer.String()]; !ok {
				return []int{}
			} else if val > 1 {
				mp[buffer.String()]--
			} else {
				delete(mp, buffer.String())
			}
			wc--
			counter = 0
			buffer.Reset()
		}
	}

	if wc == 0 {
		return []int{idx}
	}

	return []int{}
}
