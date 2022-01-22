package lc20

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsvalid(t *testing.T) {
	Convey("give ()", t, func() {
		input := "()"
		Convey("should be true", func() {
			So(isValid(input), ShouldBeTrue)
		})
	})
	Convey("give ()[]{}", t, func() {
		input := "()[]{}"
		Convey("should be true", func() {
			So(isValid(input), ShouldBeTrue)
		})
	})
	Convey("give (]", t, func() {
		input := "(]"
		Convey("should be false", func() {
			So(isValid(input), ShouldBeFalse)
		})
	})
	Convey("give ([)]", t, func() {
		input := "([)]"
		Convey("should be false", func() {
			So(isValid(input), ShouldBeFalse)
		})
	})
}

func TestIsvalidSimple(t *testing.T) {
	Convey("give ()", t, func() {
		input := "()"
		Convey("should be true", func() {
			So(isValidSimple(input), ShouldBeTrue)
		})
	})
	Convey("give ()[]{}", t, func() {
		input := "()[]{}"
		Convey("should be true", func() {
			So(isValidSimple(input), ShouldBeTrue)
		})
	})
	Convey("give (]", t, func() {
		input := "(]"
		Convey("should be false", func() {
			So(isValidSimple(input), ShouldBeFalse)
		})
	})
	Convey("give ([)]", t, func() {
		input := "([)]"
		Convey("should be false", func() {
			So(isValidSimple(input), ShouldBeFalse)
		})
	})
}
