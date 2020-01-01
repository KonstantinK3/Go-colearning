package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/KonstantinK3/Go-colearning/hw_08/mult"
)

func main() {

	fns := []func() error{}

	for i := 0; i < 100; i++ {
		fToAdd := func() error {
			errorOrNot := rand.Intn(3)      // rerurns error if 0, [o,n)
			timeToSleep := rand.Int31n(500) // time to sleep, [o,n)
			fmt.Printf("i will sleep %v millisecs and I will return error: %v\n", timeToSleep, errorOrNot == 0)
			time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
			if errorOrNot == 0 {
				return errors.New("some error here")
			}
			return nil
		}

		fns = append(fns, fToAdd)
	}

	mult.Mult(fns, 10, 3)

}
