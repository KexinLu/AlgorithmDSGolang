package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuickSort(t *testing.T) {
	Convey("given input 3,2,4,5,1", t, func() {
		input := []int{3, 2, 4, 5, 1}
		Convey("should produce 1,2,3,4,5", func() {
			exp := []int{1, 2, 3, 4, 5}

			So(quickSort(input), ShouldResemble, exp)
		})
	})
}
