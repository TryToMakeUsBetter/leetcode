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
		rec.init(10.0, 10.0)
		cir.init(10.0)
		convey.ShouldEqual(rec.Draw(), rec.height*rec.length)
		convey.ShouldEqual(cir.Draw(), math.Pi*cir.radius*cir.radius)
	})
}
