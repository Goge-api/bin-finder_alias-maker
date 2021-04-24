package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func FileWalk(p string) {
	files, err := ioutil.ReadDir(p)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f)
		if f.IsDir() {
			FileWalk(f.Name())
		}
	}
}
