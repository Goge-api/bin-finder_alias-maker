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
		println("Files Indexed", county)
		files = append(files, path)
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
			FileWritter(file)
		}
	}
}

func FileWalk(fileSearchingFor string) {
	// var wg sync.WaitGroup

	// TODO make gorouitine

	// wg.Add(1)
	// TODO make goroutine in finder await after slice is made

	// wg.Wait()

	print("benis")

}
