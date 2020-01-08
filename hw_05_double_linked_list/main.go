package main

import (
	"fmt"

	"github.com/KonstantinK3/Go-colearning/hw_05/dllist"
)

func main() {
	list := dllist.NewList()
	list.PushFront(3.14)
	list.PushBack("Jigurda")
	list.PushFront(100500)
	fmt.Printf("%v %v \n", list, list.Len())

	currentNode := list.First()
	list.Remove(currentNode)
	fmt.Printf("%v %v \n", list, list.Len())

	currentNode = list.Last()
	list.Remove(currentNode)
	fmt.Printf("%v %v \n", list, list.Len())

	list.PushFront(1)
	list.PushBack(2)
	fmt.Printf("%v %v \n", list, list.Len())

	currentNode = list.First()
	nextNode := currentNode.Next()
	list.Remove(nextNode)
	fmt.Printf("%v %v \n", list, list.Len())

	firstNode := list.First()
	fmt.Printf("value of first node is %v \n", firstNode.ValueOf())
}

// $ go run main.go
// [Head] -> [100500] -> [3.14] -> [Jigurda] -> [Tail] 3
// [Head] -> [3.14] -> [Jigurda] -> [Tail] 2
// [Head] -> [3.14] -> [Tail] 1
// [Head] -> [1] -> [3.14] -> [2] -> [Tail] 3
// [Head] -> [1] -> [2] -> [Tail] 2
// value of first node is 1
