package design

import (
	"math"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDraw(t *testing.T) {
	convey.Convey("factory", t, func() {
		var rec Rectangle
		var cir Circle
		var sh1 Shape
		var sh2 Shape
		var guestRec Rectangle
		sh1 = &guestRec
		guestCir := new(Circle)
		guestCir.Init(5.0)
		sh2 = guestCir

		guestRec.Init(5.0, 5.0)
		rec.Init(10.0, 10.0)
		cir.Init(10.0)
		convey.So(rec.Draw(), convey.ShouldEqual, 100.00)
		convey.So(cir.Draw(), convey.ShouldEqual, math.Pi*cir.radius*cir.radius)
		convey.So(sh1.Draw(), convey.ShouldEqual, guestRec.height*guestRec.length)
		convey.So(sh2.Draw(), convey.ShouldEqual, math.Pi*(*guestCir).radius*(*guestCir).radius)
	})
}
