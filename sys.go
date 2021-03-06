package main

import (
	"fmt"
	"os/exec"
)

var execCommand = exec.Command

func AliasMaker(path string, aliasWanted string) (alias string) {

	alias = aliasWanted + "=" + "'" + path + "'"
	return
}

func Execute(thing string) (string, error) {
	// print(thing)
	cmd := execCommand(thing)
	out, err := cmd.CombinedOutput()
	// fmt.Printf(string(out))
	return string(out), err
}

func SysBeep() (string, error) {
	cmd := execCommand("echo", "-ne", "'\007'")
	out, err := cmd.CombinedOutput()
	fmt.Printf(string(out))
	return string(out), err
}
