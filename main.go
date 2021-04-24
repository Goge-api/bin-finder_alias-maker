package main

import (
	"fmt"
	"sys"
)

func main() {

	go sys.Execute("1st goroutine")
	go sys.Execute("2nd goroutine")
	sys.Execute("Normal sync process")
	fmt.Println("donzo")
}
