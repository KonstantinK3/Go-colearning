package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timeServer := "0.beevik-ntp.pool.ntp.org"
	nTime, err := ntp.Time(timeServer)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Oh my, something bad happened: %d \n", err)
	}

	fmt.Printf("Current time from server %s is: %s \n", timeServer, nTime)
	fmt.Printf("Local machine time is: %s \n", time.Now())
}
