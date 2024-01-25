package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, err chan<- error, ctrl chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		select {
		case <-ctrl:
			fmt.Printf("Worker %d canceled\n", id)
			return
		default:
			fmt.Printf("Worker %d processing job %d\n", id, job)
			result := processJob(job)
			if result < 0 {
				fmt.Printf("Error in Worker %d, job %d, canceling all workers\n", id, job)
				err <- errors.New("ada error")
				close(ctrl)
				return
			}

			results <- result
		}
	}
}

func processJob(job int) int {
	if job == 2 {
		fmt.Printf("Error in job %d\n", job)
		return -1
	}

	return job * 2
}

func main() {
	const numJobs = 100

	var (
		jobs       = make(chan int, numJobs)
		results    = make(chan int, numJobs)
		err        = make(chan error, 1)
		ctrl       = make(chan struct{})
		wg         sync.WaitGroup
		numWorkers = int(math.Ceil(float64(numJobs) / 50.0))
	)

	//--- Run Routine
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, err, ctrl, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
		close(err)
	}()

	//--- Collect Error
	for itemErr := range err {
		fmt.Printf("Error -> %v\n", itemErr)
		return
	}

	//--- Collect Results
	for result := range results {
		fmt.Printf("Received Result %d\n", result)
	}
}
