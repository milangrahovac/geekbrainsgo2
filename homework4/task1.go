package main

import (
	"fmt"
)

func main() {
	var limitRoutine int = 5
	var workers = make(chan struct{}, limitRoutine)
	var m uint64 = 0
	//var mx = new(sync.Mutex)
	for i := 1; i <= 100; i++ {
		workers <- struct{}{}
		fmt.Println(i)
		go func() {
			defer func() {
				<-workers
			}()
			//	mx.Lock()
			m++
			//	mx.Unlock()
		}()
	}

	// Waiting for goroutines to finish
	for i := 0; i < limitRoutine; i++ {
		workers <- struct{}{}
	}
	fmt.Println("m:", m)
}
