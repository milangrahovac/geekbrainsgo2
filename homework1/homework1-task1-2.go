package main

import (
	"fmt"
	"log"
)

func main() {
	s := [5]int{}

	defer func() {
		if v := recover(); v != nil {
			log.Println("panic occurred:", v)
		}
	}()

	for x := 0; x < 10; x += 1 {
		fmt.Println(x, s[x])
	}

	// #not reachable
	fmt.Println("lol")

}
