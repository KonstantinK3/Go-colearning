package main

import (
	"fmt"
	"os"
)

func main() {
	var varE1, varE2 string

	varE1, ok := os.LookupEnv("E1")
	if !ok {
		fmt.Println("There is no E1")
	} else {
		fmt.Printf("E1 is equal %v\n", varE1)
	}

	varE2, ok = os.LookupEnv("E2")
	if !ok {
		fmt.Println("There is no E2")
	} else {
		fmt.Printf("E2 is equal %v\n", varE2)
	}

}
