package mult

import (
	"sync"
)

// Mult
func Mult(fns []func() error, nTasks int, nErrors int) {

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
