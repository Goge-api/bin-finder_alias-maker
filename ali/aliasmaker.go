package main

import (
	"fmt"
	"os/exec"
)

var execCommand = exec.Command

func ExecuteAlias(binAli string) (string, error) {
	// userCmd := "alias"
	// print(binAli)
	// var ali string = binAli
	// print(ali)
	cmd := execCommand("echo", "goge")
	out, err := cmd.CombinedOutput()
	fmt.Printf(string(out))
	return string(out), err
}

func main() {
	ExecuteAlias("bob")
}
