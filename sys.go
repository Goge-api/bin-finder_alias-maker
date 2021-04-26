package main

import (
	"fmt"
	"os/exec"
)

func AliasMaker(path string, aliasWanted string) (alias string) {

	alias = aliasWanted + "=" + path
	return
}

func ExecuteAlias(binAli string) (string, error) {
	// userCmd := "alias"
	var ali string = binAli
	var execCommand = exec.Command
	cmd := execCommand("alias", ali)
	out, err := cmd.CombinedOutput()
	fmt.Printf(string(out))
	return string(out), err
}
