package lc7

import "math"

/**
https://leetcode.com/problems/sequential-digits/
**/

func sequentialDigits(low int, high int) []int {
	rst := []int{}
	num := low
	for num < high {
		r, next := build(num)
		if r != -1 && r <= high {
			rst = append(rst, r)
		}
		num = next
	}
	return rst
}

func build(in int) (int, int) {
	d := in
	dgs := 0
	for d >= 10 {
		dgs++
		d /= 10
	}
	next := (d + 1) * int(math.Pow(float64(10), float64(dgs)))
	if d+dgs > 9 {
		return -1, next
	}
	built := d
	for dgs > 0 {
		built = built*10 + d + 1
		d = d + 1
		dgs--

	}

	if built < in {
		return -1, next
	}
	return built, next
}
