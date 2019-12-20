package main

import (
	"fmt"

	"github.com/KonstantinK3/Go-colearning/hw_05/dllist"
)

func main() {
	list := dllist.NewList()

	list.PushFront(1.15)
	list.PushBack(100500)

	firstNode := list.PushFront(1.15)
	lastNode := list.PushBack(100500)
	list.PushBefore(lastNode, "Jigurda")
	list.PushAfter(firstNode, "Вася")

	tmpNode := list.PushBefore(firstNode, nil)
	list.Remove(tmpNode)

	fmt.Printf("Doubly Linked List: %s\n", list)
}
