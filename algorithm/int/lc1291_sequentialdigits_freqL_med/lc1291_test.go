package lc1291

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSequentialDigits(t *testing.T) {
	Convey("give 100,300", t, func() {
		low, high := 100, 300
		Convey("should produce [123, 234]", func() {
			exp := []int{123, 234}
			So(sequentialDigits(low, high), ShouldResemble, exp)
		})
	})
	Convey("give 1000,13000", t, func() {
		low, high := 1000, 13000
		Convey("should produce [1234,2345,3456,4567,5678,6789, 12345]", func() {
			exp := []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}
			So(sequentialDigits(low, high), ShouldResemble, exp)
		})
	})
}
