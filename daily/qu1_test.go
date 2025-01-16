package daily

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	convey.Convey("test", t, func() {
		res0 := MergeTwoSortedArrays([]int{1, 4, 8}, []int{2, 6, 9})
		convey.So(res0, convey.ShouldResemble, []int{1, 2, 4, 6, 8, 9})
	})
}
