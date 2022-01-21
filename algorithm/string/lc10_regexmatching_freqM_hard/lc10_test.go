package lc10

import (
	"testing"

	//. "github.com/onsi/gomega"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIsMatch(t *testing.T) {
	Convey("Given input string aaaaa", t, func() {
		str := "aaaaa"
		Convey("When pattern is a", func() {
			p := "a"
			Convey("should not match", func() {
				So(isMatch(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is a*", func() {
			p := "a*"
			Convey("should match", func() {
				So(isMatch(str, p), ShouldBeTrue)
			})

		})
		Convey("When pattern is ab*", func() {
			p := "ab*"
			Convey("should not match", func() {
				So(isMatch(str, p), ShouldBeFalse)
			})

		})
	})
	Convey("Given input aa", t, func() {
		str := "aa"
		Convey("When pattern is a", func() {
			p := "a"
			Convey("should not match", func() {
				So(isMatch(str, p), ShouldBeFalse)
			})

		})
	})

	Convey("Given input string aabb", t, func() {
		str := "aabb"
		Convey("When pattern is a*", func() {
			p := "a*"
			Convey("should not match", func() {
				So(isMatch(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is ab*", func() {
			p := "ab*"
			Convey("should not match", func() {
				So(isMatch(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is a*b*", func() {
			p := "a*b*"
			Convey("should match", func() {
				So(isMatch(str, p), ShouldBeTrue)
			})

		})
	})

	Convey("Given input string aab", t, func() {
		str := "aab"
		Convey("When pattern is c*a*b", func() {
			p := "c*a*b"
			Convey("should match", func() {
				So(isMatch(str, p), ShouldBeTrue)
			})

		})
	})

	Convey("Given input string a", t, func() {
		str := "a"
		Convey("When pattern is .*..a", func() {
			p := ".*..a"
			Convey("should not match", func() {
				So(isMatch(str, p), ShouldBeFalse)
			})

		})
	})
}

func TestIsMatchMemo(t *testing.T) {
	Convey("Given input string aaaaa", t, func() {
		str := "aaaaa"
		Convey("When pattern is a", func() {
			p := "a"
			Convey("should not match", func() {
				So(isMatchMemo(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is a*", func() {
			p := "a*"
			Convey("should match", func() {
				So(isMatchMemo(str, p), ShouldBeTrue)
			})

		})
		Convey("When pattern is ab*", func() {
			p := "ab*"
			Convey("should not match", func() {
				So(isMatchMemo(str, p), ShouldBeFalse)
			})

		})
	})
	Convey("Given input aa", t, func() {
		str := "aa"
		Convey("When pattern is a", func() {
			p := "a"
			Convey("should not match", func() {
				So(isMatchMemo(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is a*", func() {
			p := "a*"
			Convey("should match", func() {
				So(isMatchMemo(str, p), ShouldBeTrue)
			})

		})
	})

	Convey("Given input string aabb", t, func() {
		str := "aabb"
		Convey("When pattern is a*", func() {
			p := "a*"
			Convey("should not match", func() {
				So(isMatchMemo(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is ab*", func() {
			p := "ab*"
			Convey("should not match", func() {
				So(isMatchMemo(str, p), ShouldBeFalse)
			})

		})
		Convey("When pattern is a*b*", func() {
			p := "a*b*"
			Convey("should match", func() {
				So(isMatchMemo(str, p), ShouldBeTrue)
			})

		})
	})

	Convey("Given input string aab", t, func() {
		str := "aab"
		Convey("When pattern is c*a*b", func() {
			p := "c*a*b"
			Convey("should match", func() {
				So(isMatchMemo(str, p), ShouldBeTrue)
			})

		})
	})

	Convey("Given input string a", t, func() {
		str := "a"
		Convey("When pattern is .*..a", func() {
			p := ".*..a"
			Convey("should not match", func() {
				So(isMatchMemo(str, p), ShouldBeFalse)
			})

		})
	})
}
