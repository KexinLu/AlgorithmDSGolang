package lc17

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterCombinations(t *testing.T) {
	Convey("Given 23", t, func() {
		input := "23"
		Convey("Should output ['ad','ae','af','bd','be','bf','cd','ce','cf']", func() {
			exp := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
			So(letterCombinations(input), ShouldResemble, exp)
		})
	})
	Convey("Given ", t, func() {
		input := ""
		Convey("Should output []", func() {
			exp := []string{}
			So(letterCombinations(input), ShouldResemble, exp)
		})
	})
	Convey("Given 2", t, func() {
		input := "2"
		Convey("Should output ['a','b','c']", func() {
			exp := []string{"a", "b", "c"}
			So(letterCombinations(input), ShouldResemble, exp)
		})
	})
}
