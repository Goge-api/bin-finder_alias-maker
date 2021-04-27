package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var files []string
var btc_bin_found bool = false

func GetRoot() (rootPath string) {
	root := os.Getenv("SystemDrive") + string(os.PathSeparator)
	rootPath = root
	return
}

func FilePath() (path string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return
}

func printNoRegRet(comic string, count int) {
	// println(comic, count)
	fmt.Printf("\r%d"+comic, count)

}

// func normPrint(str string) {
// 	fmt.Printf("\r%d" + str)

// }

func matcher(base, fileToFind, path string) {

	if base == fileToFind {
		FileWritter(fileToFind, path)

		fmt.Println("FOUND BITCOIN!")
		fmt.Println("Wrote paths and alii to logs")
		btc_bin_found = true
		// return nil
	} else {
		// go normPrint("...")

	}
}

func isFound() (err error) {
	if btc_bin_found == true {
		SysBeep()
		os.Exit(0)
	}
	if err != nil {
		panic(err)
	}
	return
}

func WalkAndFind(rootPath string, fileToFind string) {
	county := 0
	go printNoRegRet("Starting file search file system", county)
	root := rootPath
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		base := filepath.Base(path)
		go matcher(base, fileToFind, path)
		county++
		go printNoRegRet(" <=== This many files looked at that were not the things we are looking for.", county)
		go isFound()
		return nil
	})

	if err != nil {
		panic(err)
	}
	if err == io.EOF {
		err = nil
	}

}

// func FindParent(dir string) (parent string) {
// 	par := filepath.Dir(dir)
// 	fmt.Println(par)
// 	parent = par + `\`
// 	return
// }

// func IndexDirs(rootPath string) {

// 	county := 0
// 	root := rootPath
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		county++
// 		println("Files Indexed: ", county)
// 		files = append(files, path)
// 		return nil
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// }

// func SearchFor(fileSearchingFor string) {
// 	// Execute("echo Searching for your file")
// 	for _, file := range files {
// 		base := filepath.Base(file)
// 		fmt.Println(base)
// 		if base == fileSearchingFor {
// 			fmt.Println("FOUND BITCOIN!")
// 			// Execute(fmt.Sprintf("echo Bitcoin bin found at %s", file))
// 			FileWritter(base, file)
// 		}
// 	}
// }

// func FileWalk(rootPath string, fileSearchingFor string) {
// 	// var wg sync.WaitGroup

// 	// TODO make gorouitine
// 	go IndexDirs(rootPath)

// 	// wg.Add(1)
// 	// TODO make goroutine in finder await after slice is made
// 	go SearchFor(fileSearchingFor)

// 	// wg.Wait()

// }
