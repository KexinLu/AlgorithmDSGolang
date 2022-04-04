package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRadix(t *testing.T) {
	Convey("given input 3,1,4,6,125,2,7,5, 47, 29", t, func() {
		input := []int{3, 1, 4, 6, 125, 2, 7, 5, 47, 29}
		Convey("Should produce 1,2,3,4,5,6,7, 29, 47, 125", func() {
			expected := []int{1, 2, 3, 4, 5, 6, 7, 29, 47, 125}
			So(radix(input), ShouldResemble, expected)
		})
	})
}
