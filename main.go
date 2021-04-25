package main

import (
	"fmt"
)

func main() {

	IndexDirs(GetRoot())
	SearchFor("bitcoin.exe")

	fmt.Println("Main done running")
}
