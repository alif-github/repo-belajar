package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, number []int) {
	defer wg.Done()

	for numbers := range number {
		fmt.Printf("Worker %d starting\n", numbers)

		time.Sleep(time.Second * 3)
		fmt.Printf("Worker %d done\n", numbers)
	}
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(&wg, number)
	}

	wg.Wait()

	fmt.Println("Selesai")
}
