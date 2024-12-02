package list

import "fmt"

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
