package lc16

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSumClosest(t *testing.T) {
	Convey("Given input", t, func() {
		input := []int{-1, 2, 1, -4}
		expect := 2
		Convey("target of 1", func() {
			So(threeSumClosest(input, 1), ShouldEqual, expect)
		})
	})

	Convey("Given input", t, func() {
		input := []int{0, 0, 0, 0}
		expect := 0
		Convey("target of 1", func() {
			So(threeSumClosest(input, 1), ShouldEqual, expect)
		})
	})
}
