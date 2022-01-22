package lc12

import (
	"testing"

	//. "github.com/onsi/gomega"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIntToRoman(t *testing.T) {
	Convey("Given 3", t, func() {
		input := 3
		Convey("Should produce III", func() {
			So(intToRoman(input), ShouldEqual, "III")
		})

	})
	Convey("Given 58", t, func() {
		input := 58
		Convey("Should produce LVIII", func() {
			So(intToRoman(input), ShouldEqual, "LVIII")
		})

	})

	Convey("Given 1994", t, func() {
		input := 1994
		Convey("Should produce MCMXCIV", func() {
			So(intToRoman(input), ShouldEqual, "MCMXCIV")
		})

	})

	Convey("Given 60", t, func() {
		input := 60
		Convey("Should produce LX", func() {
			So(intToRoman(input), ShouldEqual, "LX")
		})

	})
}
