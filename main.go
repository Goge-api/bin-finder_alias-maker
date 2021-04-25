package main

import "fmt"

// func anotherThing() {
// 	go SearchFor("bitcoin.exe")
// }

func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Do you want to autorun the finder for bitcoin?")
	var auto string
	fmt.Scanln(&auto)
	if auto == "y" || auto == "Y" {
		WalkAndFind(GetRoot(), "bitcoin.exe")

	} else {

		// var then variable name then variable type

		fmt.Println("Enter the file name you're looking for.")

		// var then variable name then variable type
		var fileName string
		fmt.Scanln(&fileName)

		// Taking input from user

		// Addition of two string

		WalkAndFind(GetRoot(), fileName)
	}

	// IndexDirs(GetRoot())

	// SearchFor("bitcoin.exe")
	// anotherThing()

	println("thing done")
}
