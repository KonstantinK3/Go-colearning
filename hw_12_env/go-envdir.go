package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KonstantinK3/Go-colearning/hw_12_env/envdir"
)

func main() {

	// тестовый запуск: ./go-envdir ./tester/vars ./tester/printer
	// результат:
	// E1 is equal oopsE1
	// E2 is equal E2_variable
	// commad executed with exit status 0

	args := os.Args

	if len(args) != 3 {
		log.Fatalf("please provide exactly 2 arguments")
	}

	cmd := []string{}
	// cmd = append(cmd, "./tester/printer")
	cmd = append(cmd, args[2])

	// env, err := envdir.ReadDir("./tester/vars")
	env, err := envdir.ReadDir(args[1])

	if err != nil {
		log.Fatalf("error occured: %s\n", err)
	}

	res := envdir.RunCmd(cmd, env)
	fmt.Printf("commad executed with exit status %v \n", res)

}
