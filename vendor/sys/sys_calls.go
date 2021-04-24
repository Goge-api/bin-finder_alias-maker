package sys

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func Execute(thing string) {
	for i := 0; i < 99; i++ {
		stringy := strconv.Itoa(i)
		cmd := exec.Command("echo", thing, stringy)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(out.String())
	}

}
