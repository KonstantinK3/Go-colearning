package main

import (
	"flag"
	"log"

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
	err := copy.Copy(from, to, offset, limit)

	if err != nil {
		log.Fatalf("Oh my, something bad happened: '%s' \n", err)
	}

}
