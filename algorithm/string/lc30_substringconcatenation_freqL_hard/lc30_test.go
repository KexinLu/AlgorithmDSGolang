package lc30

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindSubstring(t *testing.T) {
	Convey("for input barfoothefoobarman and [foo, bar]", t, func() {
		input := "barfoothefoobarman"
		word := []string{"foo", "bar"}
		Convey("should produce [0,9]", func() {
			exp := []int{0, 9}
			So(findSubstring(input, word), ShouldResemble, exp)
		})
	})
	Convey("for input barfoothefoobarman and [foo, bar]", t, func() {
		input := "barfoothefoobarman"
		word := []string{"aaa", "ddd"}
		Convey("should produce []", func() {
			exp := []int{}
			So(findSubstring(input, word), ShouldResemble, exp)
		})
	})
	Convey("for input barfoothefoobarman and []", t, func() {
		input := "barfoothefoobarman"
		word := []string{}
		Convey("should produce []", func() {
			exp := []int{}
			So(findSubstring(input, word), ShouldResemble, exp)
		})
	})
	Convey("for input '' and [abc]", t, func() {
		input := ""
		word := []string{"abc"}
		Convey("should produce []", func() {
			exp := []int{}
			So(findSubstring(input, word), ShouldResemble, exp)
		})
	})
	Convey("for input barfoothefoobarman and [foo, bar]", t, func() {
		input := "rdgoodgoodgoodbestword"
		word := []string{"word", "good", "best", "word"}
		Convey("should produce []", func() {
			exp := []int{}
			So(findSubstring(input, word), ShouldResemble, exp)
		})
	})
	Convey("for input barfoofoobarthefoobarman and [foo, bar]", t, func() {
		input := "barfoofoobarthefoobarman"
		word := []string{"bar", "foo", "the"}
		Convey("should produce []", func() {
			exp := []int{6, 9, 12}
			So(findSubstring(input, word), ShouldResemble, exp)
		})
	})
}
