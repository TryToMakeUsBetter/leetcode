package list

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAddList(t *testing.T) {
	convey.Convey("NewLisr", t, func() {
		node1 := NewNode(1)
		node2 := NewNode(2)
		node3 := NewNode(3)
		list1 := NewList(node1)
		list1.HeadAddList(node2)
		list1.RearAddList(node3)

		list1.Display()
		//list2.Display()
	})

}

func TestReverseList(t *testing.T) {
	convey.Convey("ReverseList", t, func() {
		list1 := TurnSliceIntoList([]int{1, 2, 3, 4, 5})
		list1.Display()
		fmt.Println()

		list1.ReverseList()
		list1.Display()
		fmt.Println()

		list1.ReverseListv2()
		list1.Display()
		fmt.Println()

		list1.ReversePartList(2, 5)
		list1.Display()
	})
}
