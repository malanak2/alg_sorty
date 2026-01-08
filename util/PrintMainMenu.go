package util

import (
	"fmt"
	"io/fs"
)

func PrintMainMenu(files []fs.DirEntry) {
	fmt.Printf("Please choose a dataset to sort:\n")
	for i, file := range files {
		fmt.Printf("%d: %s\n", i+1, file.Name())
	}
	fmt.Printf("0: Exit\n")
}
