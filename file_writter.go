package main

import (
	"fmt"
	"log"
	"os"
)

func FileWritter(ff string, path string) {

	f, err := os.OpenFile("binsFound.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(ff + "," + path + "\n"); err != nil {
		log.Println(err)
	}

	fmt.Println("File Writter done")
}
