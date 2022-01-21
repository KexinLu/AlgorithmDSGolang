package lc134

import (
	"testing"

	//. "github.com/onsi/gomega"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCanCompleteCircuit(t *testing.T) {
	Convey("Given gas [1,2,3,4,5] ", t, func() {
		gas := []int{1, 2, 3, 4, 5}
		Convey("When cost is [3,4,5,1,2]", func() {
			cost := []int{3, 4, 5, 1, 2}
			Convey("should return 3", func() {
				So(canCompleteCircuit(gas, cost), ShouldEqual, 3)
			})

		})
	})

	Convey("Given gas [2,3,4] ", t, func() {
		gas := []int{2, 3, 4}
		Convey("When cost is [3,4,3]", func() {
			cost := []int{3, 4, 3}
			Convey("should return -1", func() {
				So(canCompleteCircuit(gas, cost), ShouldEqual, -1)
			})

		})
	})
}
