package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	fileNums = 0
	dirNums = 0
)

func listAll(path string, depth int) {
	readerInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range readerInfos {
		if info.IsDir() {
			outPut(depth)
			dirNums++
			fmt.Println(info.Name())
			listAll(filepath.Join(path, info.Name()), depth+1)
		}else {
			outPut(depth)
			fileNums++
			fmt.Println(info.Name())
		}
	}
}

func outPut(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("\t ")
	}
//	fmt.Println("|")
//	for j := 0; j < depth+1; j++ {
//		fmt.Print("___ ")
//	}
}

func main() {
	pathptr := flag.String("path", "", "dirpath to read from")
	flag.Parse()

	if pathptr == nil {
		dir, err := os.Getwd()
		if err != nil {
			os.Exit(1)
		}
		pathptr = &dir
		fmt.Println(".")
	}else {
		fmt.Println(*pathptr)
	}

	listAll(*pathptr, 1)
	fmt.Printf("%d directories, %d files\n", dirNums, fileNums)
}
