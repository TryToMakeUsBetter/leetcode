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
		convey.ShouldEqual(rec.Draw(), rec.height*rec.length)
		convey.ShouldEqual(cir.Draw(), math.Pi*cir.radius*cir.radius)
		convey.ShouldEqual(sh1.Draw(), guestRec.height*guestRec.length)
		convey.ShouldEqual(sh2.Draw(), math.Pi*(*guestCir).radius*(*guestCir).radius)
	})
}
