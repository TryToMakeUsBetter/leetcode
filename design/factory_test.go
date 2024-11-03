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
		var guestRec Rectangle
		sh1 = &guestRec
		guestRec.Init(5.0, 5.0)
		rec.Init(10.0, 10.0)
		cir.Init(10.0)
		convey.ShouldEqual(rec.Draw(), rec.height*rec.length)
		convey.ShouldEqual(cir.Draw(), math.Pi*cir.radius*cir.radius)
		convey.ShouldEqual(sh1.Draw(), guestRec.height*guestRec.length)
	})
}
