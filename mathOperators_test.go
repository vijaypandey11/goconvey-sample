package testing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//
func TestAdd(t *testing.T) {

	Convey("The result of 5 + 11 = 16 ?", t, func() {
		x, y := 5, 11
		So(Add(x, y), ShouldEqual, 16)
	})
}

func TestSubtract(t *testing.T) {
	Convey("The result of 24 - 11 != 22 ?", t, func() {
		x, y := 24, 11
		So(Subtract(x, y), ShouldNotEqual, 22)
	})
}

func TestMultiply(t *testing.T) {
	Convey("The result of 11 * 2 = 22 ?", t, func() {
		x, y := 11, 2
		So(Multiply(x, y), ShouldEqual, 22)
	})
}

func TestDivision(t *testing.T) {
	Convey("division of even numbers", t, func() {
		x, y := 10, 2
		Convey("Divide by non-zero number", func() {
			n, _ := Division(x, y)
			So(n, ShouldEqual, 5)
		})

		Convey("Divide by zero", func() {
			y = 0
			_, err := Division(x, y)
			So(err, ShouldNotBeNil)
		})
	})
}
