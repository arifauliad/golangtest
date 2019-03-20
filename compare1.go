package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	fileInfoSource os.FileInfo
	errSource      error
)

var (
	fileInfoTarget os.FileInfo
	errTarget      error
)

func main() {
	var sourcefiles []string
	var targetfiles []string
	var is_ int

	rootSource := "source"
	errSource := filepath.Walk(rootSource, func(path string, info os.FileInfo, err error) error {
		sourcefiles = append(sourcefiles, path)
		return nil
	})
	if errSource != nil {
		panic(errSource)
	}

	rootTarget := "target"
	errTarget := filepath.Walk(rootTarget, func(path string, info os.FileInfo, err error) error {
		targetfiles = append(targetfiles, path)
		return nil
	})
	if errTarget != nil {
		panic(errTarget)
	}

	for _, sourcefile := range sourcefiles {
		fileInfoSource, errSource = os.Stat(sourcefile)
		is_ = 0
		if fileInfoSource.IsDir() == false {
			for _, targetfile := range targetfiles {
				fileInfoTarget, errTarget = os.Stat(targetfile)
				if fileInfoSource.Name() == fileInfoTarget.Name() && fileInfoSource.Size() == fileInfoTarget.Size() {
					is_ = 1
				}
			}
			if is_ == 0 {
				fmt.Println(sourcefile, " NEW")
			}
		}
	}

	for _, targetfile := range targetfiles {
		fileInfoTarget, errTarget = os.Stat(targetfile)
		is_ = 0
		if fileInfoTarget.IsDir() == false {
			for _, sourcefile := range sourcefiles {
				fileInfoSource, errSource = os.Stat(sourcefile)
				if fileInfoTarget.Name() == fileInfoSource.Name() && fileInfoTarget.Size() == fileInfoSource.Size() {
					is_ = 1
				}
			}
			if is_ == 0 {
				fmt.Println(targetfile, " DELETED")
			}
		}
	}
}
