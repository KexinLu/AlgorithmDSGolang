package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSelectionSort(t *testing.T) {
	Convey("Given input 3, 1, 4, 5, 6,2", t, func() {
		input := []int{3, 1, 4, 5, 6, 2}
		Convey("Should produce 1,2,3,4,5,6", func() {
			expected := []int{1, 2, 3, 4, 5, 6}
			So(selectionSort(input), ShouldResemble, expected)
		})
	})

}
