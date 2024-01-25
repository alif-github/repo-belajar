package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second * 3)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		a := i

		go func() {
			defer wg.Done()
			worker(a)
		}()
	}

	wg.Wait()

	fmt.Println("Selesai")

}
