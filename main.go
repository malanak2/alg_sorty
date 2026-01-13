package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/malanak2/alg_sorts/sorts"
	"github.com/malanak2/alg_sorts/util"
	"golang.org/x/term"
)

func main() {
	data_dir := "./data"

	files, err := os.ReadDir(data_dir)
	if err != nil {
		panic(err)
	}

	util.PrintMainMenu(files)

	// Read only characters without waiting for Enter key
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	term.Restore(int(os.Stdin.Fd()), oldState)

	if rune(b[0]) > rune(strconv.Itoa(len(files))[0]) || rune(b[0]) < rune('0') {
		fmt.Printf("\nInvalid option, exiting...\n")
		return
	}

	file := files[int(b[0]-'1')]
	if b[0] == '0' {
		fmt.Println("\nExiting...")
		return
	}

	util.PrintSortsOptions()
	oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	b = make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	term.Restore(int(os.Stdin.Fd()), oldState)

	if b[0] > '7' || b[0] < '0' {
		fmt.Println("\nInvalid option, exiting...")
		return
	}
	if b[0] == '0' {
		fmt.Println("\nExiting...")
		return
	}
	switch b[0] {
	case '1':
		sorts.SelectionSort(file)
	case '2':
		sorts.BubbleSort(file)
	case '3':
		sorts.InsertionSort(file)
	case '4':
		sorts.HeapSort(file)
	}
}
