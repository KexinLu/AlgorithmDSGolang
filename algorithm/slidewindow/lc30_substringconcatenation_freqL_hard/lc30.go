package lc30slide

func findSubstring(s string, words []string) []int {
	rst := []int{}
	if s == "" {
		return rst
	}
	if len(words) == 0 {
		return rst
	}
	wm := make(map[string]int)
	for _, w := range words {
		if _, ok := wm[w]; !ok {
			wm[w] = 1
		} else {
			wm[w]++
		}
	}
	wl := len(words[0])
	wc := len(words)
	rst = make([]int, 0)
	for i := 0; i < wl; i++ {
		rst = append(rst, walk(i, s, wl, wc, wm)...)
	}
	return rst
}

type Q []string

func (q *Q) Pop() string {
	cur := *q
	out := cur[0]
	*q = cur[1:]
	return out
}

func (q *Q) Push(in string) {
	*q = append(*q, in)
}

func (q *Q) Clear() {
	*q = []string{}
}

func walk(start int, s string, wl int, wc int, wm map[string]int) []int {
	rst := []int{}
	if len(s) < wl*wc {
		return rst
	}

	mp := make(map[string]int)
	popMap := func() {
		for k, v := range wm {
			mp[k] = v
		}
	}
	popMap()
	q := &Q{}
	i := start
	for j := i + wl; j <= len(s); j += wl {
		cur := s[j-wl : j]
		if val, ok := mp[cur]; !ok {
			i = j
			q.Clear()
			popMap()
		} else {
			q.Push(cur)
			mp[cur]--
			for val < 0 {
				first := q.Pop()
				mp[first]++
				i += wl
				val, _ = mp[cur]
			}
			if len(*q) == wc {
				shouldAppend := true
				for _, v := range mp {
					if v != 0 {
						shouldAppend = false
					}
				}
				if shouldAppend {
					rst = append(rst, i)
				}
				first := q.Pop()
				mp[first]++
				i += wl
			}
		}
	}
	return rst
}
