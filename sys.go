package main

import (
	"fmt"
	"os/exec"
)

var execCommand = exec.Command

func AliasMaker(path string, aliasWanted string) (alias string) {

	alias = `"` + aliasWanted + "=" + "'" + path + "'" + `"`
	return
}

func ExecuteAlias(binAli string) (string, error) {
	// userCmd := "alias"
	// print(binAli)
	var ali string = binAli
	// print(ali)
	cmd := execCommand("alias", ali)
	out, err := cmd.CombinedOutput()
	fmt.Printf(string(out))
	return string(out), err
}
func Execute(thing string) (string, error) {
	print(thing)
	cmd := execCommand(thing)
	out, err := cmd.CombinedOutput()
	fmt.Printf(string(out))
	return string(out), err
}
