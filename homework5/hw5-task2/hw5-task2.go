// Реализуйте функцию для разблокировки мьютекса с помощью defer

package main

import (
	"fmt"
	"sync"
)

var s int = 10

func counter(m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	s += 1
}

func main() {
	fmt.Println("task 2")
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func() {
			counter(&mutex)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("sum:", s)
}
