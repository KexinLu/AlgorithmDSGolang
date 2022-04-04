package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBubble(t *testing.T) {

	Convey("given input 3,7,4,5,1,2", t, func() {
		input := []int{3, 7, 4, 5, 1, 2}
		Convey("should produce 1,2,3,4,5,7", func() {
			expected := []int{1, 2, 3, 4, 5, 7}
			So(bubble(input), ShouldResemble, expected)
		})
	})
}
