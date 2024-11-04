package design

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDraw(t *testing.T) {
	convey.Convey("factory", t, func() {
		sf0 := new(ShapeFactoryStruct)
		c, _ := sf0.CreateShape("Circle")
		sf1 := new(ShapeFactoryStruct)
		r, _ := sf1.CreateShape("Rectangle")
		convey.So(c, convey.ShouldHaveSameTypeAs, new(Circle))
		convey.So(r, convey.ShouldHaveSameTypeAs, new(Rectangle))
	})
}
