package interview

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCheckSubString(t *testing.T) {
	convey.Convey("qu4", t, func() {
		res := CheckSubString("ADOLBEAC", "ABC")
		convey.So(res, convey.ShouldEqual, "BEAC")
	})
}
