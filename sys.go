package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func AliasMaker(path string, aliasWanted string) (alias string) {

	alias = aliasWanted + "=" + path
	return
}

func ExecuteAlias(thing string) (string, error) {
	// userCmd := "alias"
	var ali string = BinLogRead()[0].fullAlias
	// // argument := thing
	// cmd := exec.Command("alias", ali)

	// // Run firefox-bin and load URL specified in argument.
	// err := cmd.Run()
	// fmt.Println("DONE")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	print(ali)
	var execCommand = exec.Command
	cmd := execCommand("alias", ali)
	cmd.Stdin = strings.NewReader(thing)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf(string(out))
		return string(out), err
	}
	return string(out), nil

	// Print the output
	// fmt.Println(string(stdout))

}
