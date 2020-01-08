package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KonstantinK3/Go-colearning/hw_10/copy"
)

var from string
var to string
var limit int
var offset int

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.IntVar(&offset, "offset", 0, "offset in input file")
	flag.IntVar(&limit, "limit", 0, "limit to copy")
}

func main() {

	flag.Parse()

	fmt.Printf("%T, %#v \n", os.Args, os.Args)
	fmt.Println(from, to, offset, limit)

	copy.Copy(from, to, offset, limit)

}
