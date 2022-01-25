package lc29

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDivide(t *testing.T) {
	Convey("Given dividend is 1200, divisor is 7", t, func() {
		Convey("Should give correct result", func() {
			So(divide(1200, 7), ShouldEqual, 1200/7)
		})
	})
	Convey("Given dividend is 0, divisor is 7", t, func() {
		Convey("Should give correct result", func() {
			So(divide(0, 7), ShouldEqual, 0/7)
		})
	})
	Convey("Given dividend is -21, divisor is -6 ", t, func() {
		Convey("Should give correct result", func() {
			So(divide(-21, -6), ShouldEqual, -21/-6)
		})
	})

	Convey("Given dividend is -21, divisor is 6 ", t, func() {
		Convey("Should give correct result", func() {
			So(divide(-21, 6), ShouldEqual, -21/6)
		})
	})

}
