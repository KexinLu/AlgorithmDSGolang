package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeSort(t *testing.T) {
	Convey("Given input 5, 2, 6, 1, 3, 4", t, func() {
		input := []int{5, 2, 6, 1, 3, 4}
		Convey("Should produce 1,2,3,4,5,6", func() {
			exp := []int{1, 2, 3, 4, 5, 6}
			So(mergeSort(input), ShouldResemble, exp)
		})
	})
}
