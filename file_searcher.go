package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var files []string
var btc_bin_found bool = false

func GetRoot() (ok string) {
	benis := os.Getenv("SystemDrive") + string(os.PathSeparator)
	ok = benis
	println(benis)
	return
}

func FindParent(dir string) (parent string) {
	par := filepath.Dir(dir)
	fmt.Println(par)
	parent = par + `\`
	return
}

func FilePath() (path string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return
}

func IndexDirs(rootPath string) {

	county := 0
	// Execute("echo Indexing your files")
	print("Indexing file system")
	root := rootPath
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		county++
		println("Files Indexed: ", county)
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

}

func WalkAndFind(rootPath string, thing string) {

	county := 0
	print("Starting file search file system")
	root := rootPath
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		base := filepath.Base(path)
		if base == thing {
			fmt.Println("FOUND BITCOIN!")
			FileWritter(thing, path)
			// return nil
			os.Exit(0)
		} else {

			county++
			println("Files not what we are looking for...", county)
			// files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}

func SearchFor(fileSearchingFor string) {
	// Execute("echo Searching for your file")
	for _, file := range files {
		base := filepath.Base(file)
		fmt.Println(base)
		if base == fileSearchingFor {
			fmt.Println("FOUND BITCOIN!")
			// Execute(fmt.Sprintf("echo Bitcoin bin found at %s", file))
			FileWritter(base, file)
		}
	}
}

func FileWalk(rootPath string, fileSearchingFor string) {
	// var wg sync.WaitGroup

	// TODO make gorouitine
	go IndexDirs(rootPath)

	// wg.Add(1)
	// TODO make goroutine in finder await after slice is made
	go SearchFor(fileSearchingFor)

	// wg.Wait()

	print("benis")

}
