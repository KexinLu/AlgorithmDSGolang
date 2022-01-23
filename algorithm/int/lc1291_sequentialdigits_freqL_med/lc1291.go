package lc1291

import "math"

func sequentialDigits(low int, high int) []int {
	rst := []int{}
	num := low

	for num < high {
		r, next := build(num)

		// when -1 there's no possible value in current range
		if r != -1 && r <= high {
			rst = append(rst, r)
		}
		num = next
	}
	return rst
}

func build(in int) (int, int) {
	// first digit of current number
	d := in

	// how many digits are there in this number
	dgs := 0

	// divid the number by 10 each loop and increment dgs by 1, until we are less than 10
	for d >= 10 {
		dgs++
		d /= 10
	}

	// produce the next number we want to build upon
	// if current is 1234, next will be 2000
	// when we are at 9000, next will be 10000
	// this relies on the fact that there's only one possible number fulfill the requirement between a range like 20000 and 30000, thus each time we only have to increment the first digit
	next := (d + 1) * int(math.Pow(float64(10), float64(dgs)))

	// if first digit + number of digits is greater than 9, it's wrong
	if d+dgs > 9 {
		return -1, next
	}

	// building the number
	built := d
	for dgs > 0 {
		built = built*10 + d + 1
		d = d + 1
		dgs--
	}

	// this is dealing with initial case where it's possible for us to have 1920 while built is 1234 which is less than input
	if built < in {
		return -1, next
	}

	return built, next
}
