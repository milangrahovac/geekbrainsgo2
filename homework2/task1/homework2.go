package main

import (
	"fmt"

	"github.com/matishsiao/goInfo"
)

func main() {
	fmt.Println("Your OS:")
	gi := goInfo.GetInfo()
	gi.VarDump()
}
