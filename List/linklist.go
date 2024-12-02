package list

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func NewNode(v int) *Node {
	return &Node{
		v,
		nil,
	}
}

type List struct {
	Head *Node
}

func NewList(node *Node) (list *List) {
	list = new(List)
	list.Head = node
	return list
}

func (list *List) HeadAddList(node *Node) {
	node.next = list.Head
	list.Head = node
}

func (list *List) RearAddList(node *Node) {
	current := list.Head
	for current.next != nil {
		current = current.next
	}

	rear := current
	(*rear).next = node
}

func (list List) Display() {
	current := list.Head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func (list *List) ReverseList() {
	if list.Head == nil {
		return
	}
	current := list.Head
	var pre *Node

	for current != nil {
		nxt := current.next
		current.next = pre
		pre = current
		current = nxt
	}

	list.Head = pre
}

func TurnSliceIntoList(input []int) (list *List) {
	if len(input) == 0 {
		return
	}
	h := NewNode(input[0])
	list = NewList(h)

	for i := 1; i < len(input); i++ {
		tmp := NewNode(input[i])
		list.HeadAddList(tmp)
	}

	return list
}
func (list *List) ReversePartList(m int, n int) {

	mark := &Node{
		-1,
		list.Head,
	}

	pre := mark

	for i := 0; i < m-1; i++ {
		pre = pre.next
	}
	current := pre.next

	for i := m; i < n; i++ {
		nxt := current.next
		current.next = nxt.next
		nxt.next = pre.next
		pre.next = nxt
	}

	list.Head = mark.next
}
