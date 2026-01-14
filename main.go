package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
	if b[0] > '8' || b[0] < '0' {
		fmt.Println("\nInvalid option, exiting...")
		return
	}
	if b[0] == '0' {
		fmt.Println("\nExiting...")
		return
	}
	fmt.Printf("\nYou selected file: %s\n", file.Name())
	var timeItTook time.Duration
	switch b[0] {
	case '1':
		timeItTook = sorts.SelectionSort(file, true)
	case '2':
		timeItTook = sorts.BubbleSort(file, true)
	case '3':
		timeItTook = sorts.InsertionSort(file, true)
	case '4':
		timeItTook = sorts.HeapSort(file, true)
	case '5':
		timeItTook = sorts.MergeSort(file, true)
	case '6':
		timeItTook = sorts.QuickSort(file, true)
	case '7':
		timeItTook = sorts.RadixSort(file, true)
	case '8':
		timeSelection := make(chan time.Duration, 1)
		timeBubble := make(chan time.Duration, 1)
		timeInsertion := make(chan time.Duration, 1)
		timeHeap := make(chan time.Duration, 1)
		timeMerge := make(chan time.Duration, 1)
		timeQuick := make(chan time.Duration, 1)
		timeRadix := make(chan time.Duration, 1)
		go util.SortAsync(sorts.SelectionSort, file, false, timeSelection)
		go util.SortAsync(sorts.BubbleSort, file, false, timeBubble)
		go util.SortAsync(sorts.InsertionSort, file, false, timeInsertion)
		go util.SortAsync(sorts.HeapSort, file, false, timeHeap)
		go util.SortAsync(sorts.MergeSort, file, false, timeMerge)
		go util.SortAsync(sorts.QuickSort, file, false, timeQuick)
		go util.SortAsync(sorts.RadixSort, file, false, timeRadix)

		fmt.Printf("Radix sort: %s\n", <-timeRadix)
		fmt.Printf("Quick sort: %s\n", <-timeQuick)
		fmt.Printf("Merge sort: %s\n", <-timeMerge)
		fmt.Printf("Heap sort: %s\n", <-timeHeap)
		fmt.Printf("Insertion sort: %s\n", <-timeInsertion)
		fmt.Printf("Selection sort: %s\n", <-timeSelection)
		fmt.Printf("Bubble sort: %s\n", <-timeBubble)

		fmt.Printf("\nAll sorts completed!\n")
	}
	fmt.Printf("\nSorting took %s\n", timeItTook)
}
