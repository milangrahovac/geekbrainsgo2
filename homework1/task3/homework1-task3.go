package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	for x := 0; x < 1000000; x += 1 {
		fileName := fmt.Sprintf("test/file%d.txt", x)
		emptyFile, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		emptyFile.Close()
	}

}
