package dllist

import "fmt"

// Node is a single node for list
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
	list  *List
}

// List is a double linked list of Nodes
type List struct {
	head   *Node
	tail   *Node
	length int
}

//NewList makes a new list. myList := dllist.NewList()
func NewList() *List {
	return &List{}
}

func (list *List) Len() int {
	return list.length
}

func (list *List) First() *Node {
	return list.head
}

func (list *List) Last() *Node {
	return list.tail
}

func (list *List) PushFront(value interface{}) {
	var node *Node
	if list.head == nil {
		node = &Node{value, nil, nil, list}
		list.tail = node
		list.head = node
		list.length++
	} else {
		node = &Node{value, nil, list.head, list}
		list.head.prev = node
		list.head = node
		list.length++
	}
}

func (list *List) PushBack(value interface{}) {
	var node *Node
	if list.tail == nil {
		node = &Node{value, nil, nil, list}
		list.tail = node
		list.head = node
		list.length++
	} else {
		node = &Node{value, list.tail, nil, list}
		list.tail.next = node
		list.tail = node
		list.length++
	}
}

func (lst *List) Remove(node *Node) {
	if node.list != lst {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lst.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lst.tail = node.prev
	}

	lst.length--
}

func (list *List) String() string {
	result := fmt.Sprintf("[Head] -> ")
	cursor := list.head
	for cursor != nil {
		result += fmt.Sprintf("[%v] -> ", cursor.Value)
		cursor = cursor.next
	}
	result += fmt.Sprintf("[Tail]")

	return result
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) ValueOf() interface{} {
	return node.Value
}
