package mult

import (
	"fmt"
	"sync"
)

// Mult
func Mult(fns []func() error, nTasks int, nErrors int) {

	jobs := make(chan func() error, nTasks)
	done := make(chan struct{})
	errorCounter := &ErrorCounter{}

	for i := 0; i < nTasks; i++ {
		go func(counter *ErrorCounter, idx int) {
			fmt.Printf("job %v started \n", idx)
			for {
				select {
				case fn := <-jobs:
					fmt.Printf("%v Start processing of new job\n", idx)
					if err := fn(); err != nil {
						counter.Increase()
						fmt.Printf("%v Failed to process job\n", idx)
					} else {
						fmt.Printf("%v Successfull done\n", idx)
					}
				case <-done:
					fmt.Printf("%v Got exit signal\n", idx)
					return
				}
			}
		}(errorCounter, i)
	}

	for _, job := range fns {
		if errorCounter.Value() >= nErrors {
			fmt.Printf("----------Maximum errors limit reached: %v----------\n", errorCounter.Value())
			done <- struct{}{}
			break
		}

		fmt.Println("<-- Sent job to worker")
		jobs <- job
	}

	fmt.Println("Done!")

}

func Mult_old(fns []func() error, nTasks int, nErrors int) {

	currentErrors := 0
	var wg sync.WaitGroup

	var ch = make(chan int)

	for _, v := range fns {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			err := v()
			if err != nil {
				currentErrors++
			}
			if currentErrors >= nErrors {
				ch <- 1
			}

			wg.Done()
		}(&wg)
	}

	<-ch

	wg.Wait()

}
