package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountingSort(t *testing.T) {
	Convey("input of []int{1,5,4,2,3,6}", t, func() {
		input := []int{1, 5, 4, 2, 3, 6}
		Convey("should produce 1,2,3,4,5,6", func() {
			expected := []int{1, 2, 3, 4, 5, 6}
			So(countingSort(input), ShouldResemble, expected)
		})
	})
}
