package design

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestInit(t *testing.T) {
	convey.Convey("Singleton", t, func() {
		OurSun := new(NKSun)
		(*OurSun).Set("JZE")
		(*OurSun).Set("123")
		convey.So((*OurSun).Get(), convey.ShouldEqual, "JZE")
	})
}
