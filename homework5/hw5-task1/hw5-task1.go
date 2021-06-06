// Напишите программу, которая запускает n потоков и дожидается завершения их всех

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("task 3")

	var (
		n       int = 100
		counter int = 0
		wg          = sync.WaitGroup{}
	)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n, counter)

}
