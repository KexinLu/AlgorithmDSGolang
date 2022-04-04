package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsert(t *testing.T) {
	Convey("given input []int{3,1,5,4,6,2}", t, func() {
		input := []int{3, 1, 5, 4, 6, 2}
		Convey("should produce [1,2,3,4,5,6]", func() {
			expected := []int{1, 2, 3, 4, 5, 6}
			So(expected, ShouldResemble, insert(input))
		})
	})
}
