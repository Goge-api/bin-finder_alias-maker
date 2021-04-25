package main

import (
	"fmt"
	"log"
	"os"
)

func FileWritter(s string) {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(s)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("File Writter done")
}
