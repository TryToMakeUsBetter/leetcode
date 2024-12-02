package list

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAddList(t *testing.T) {
	convey.Convey("test", t, func() {
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
