package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Welcome")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("Entrance")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 5; i++ {
			fmt.Println("Exit")
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")
	//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

}
