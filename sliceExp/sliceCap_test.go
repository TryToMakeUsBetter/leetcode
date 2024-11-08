package sliceexp

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMySlice(t *testing.T) {
	convey.Convey("slice enlarge cap", t, func() {
		res1, res2 := MySlice()
		convey.So(res1, convey.ShouldResemble, []int{0, 1, 2, 3, 4})
		convey.So(res2, convey.ShouldResemble, res1[0:2])
		EnlargeSlice1(&res1)
		convey.So(res2, convey.ShouldResemble, res1[0:2])
		convey.So(len(res1), convey.ShouldEqual, 7)
		EnlargeSliceAgain(&res1)
		convey.So(res2, convey.ShouldResemble, res1[0:2])
	})

}
