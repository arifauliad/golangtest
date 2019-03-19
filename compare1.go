package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func visitsource(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func visittarget(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func main() {
	var sourcefiles []string
	var targetfiles []string

	source := "C:/Go/src/compare1/source"
	errsource := filepath.Walk(source, visitsource(&sourcefiles))
	if errsource != nil {
		panic(errsource)
	}

	target := "C:/Go/src/compare1/target"
	errtarget := filepath.Walk(target, visittarget(&targetfiles))
	if errtarget != nil {
		panic(errtarget)
	}

	for _, file := range sourcefiles {
		fmt.Println(file)
	}

	fmt.Println()
	for _, file := range targetfiles {
		fmt.Println(file)
	}
}
