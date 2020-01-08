package mult

import (
	"fmt"
)

// Mult
func Mult(fns []func() error, nTasks int, nErrors int) {

	jobs := make(chan func() error, nTasks)
	counter := 0

	for i := 0; i < nTasks; i++ {
		go func(id int) {
			fmt.Printf("job number %v started \n", id)
			///////
			for {
				currentFunc := <-jobs
				fmt.Printf("Starting new job in goruntine number %v \n", id)
				err := currentFunc()
				if err != nil {
					counter++
					fmt.Printf("Job number %v ruturned an error '%s'\n", id, err)
				} else {
					fmt.Printf("Job number %v done succesfully\n", id)
				}

			}
		}(i)
	}

	for c, job := range fns {
		if counter >= nErrors {
			fmt.Printf("----------Maximum errors limit reached: %v----------\n", counter)
			return
		}

		fmt.Printf("=======Sending job %v to worker=======\n", c)
		jobs <- job
	}

}
