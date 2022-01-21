package lc6

import (
	"testing"

	//. "github.com/onsi/gomega"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convey("Given PAYPALISHIRING and 3 rows", t, func() {
		input := "PAYPALISHIRING"
		rows := 3
		Convey("should return PAHNAPLSIIGYIR", func() {
			So(convert(input, rows), ShouldEqual, "PAHNAPLSIIGYIR")
		})

	})

}
