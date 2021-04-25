package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func BinLogRead() (structArray []Al) {
	file, err := os.Open("binsFound.log")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	var loggins []Al

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		structVal := strings.Split(eachline, ",")
		var thing Al = Al{executable: structVal[0], path: structVal[1], aliasCmd: structVal[2], fullAlias: structVal[3]}
		loggins = append(loggins, thing)
		// fmt.Println(eachline)
	}
	structArray = loggins
	return
}
