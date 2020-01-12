package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/KonstantinK3/Go-colearning/hw_12_env/envdir"
)

func main() {

	args := os.Args

	fmt.Printf("------%#v \n", args)

	if len(args) < 3 {
		log.Fatalf("please provide at least 3 arguments: path, file and envs")
	}

}
