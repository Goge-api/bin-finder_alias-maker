package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func Execute(thing string) {

	cmd := exec.Command(thing)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(out.String())

}
