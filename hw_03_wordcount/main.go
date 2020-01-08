package main

import (
	"fmt"

	"github.com/KonstantinK3/Go-colearning/hw_03/wordcount"
)

func main() {

	fmt.Println(wordcount.Wordcount("test.txt"))
	fmt.Println(wordcount.Wordcount("voina-i-mir.txt"))

}
