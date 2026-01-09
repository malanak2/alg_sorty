package sorts

import (
	"bufio"
	"container/list"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
)

func SelectionSort(file fs.DirEntry) {
	data := list.New()
	open, err := os.Open(filepath.Join("./data", file.Name()))
	if err == nil {
		defer open.Close()
	} else {
		fmt.Println("Error opening file")
		return
	}
	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(open)

	// Loop through the file and read each line
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		lineI, err := strconv.Atoi(line)
		if err == nil {
			data.PushBack(lineI)
			continue
		}
		fmt.Printf("Non-integer value found, treating as string: %s\n", line)
		data.PushBack(line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s", err)
	}

	bar := progressbar.NewOptions64(
		int64(data.Len()),
		progressbar.OptionSetDescription(""),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(10),
		progressbar.OptionShowTotalBytes(true),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
	)
	for i := data.Front(); i != nil; i = i.Next() {
		min := i
		for j := i.Next(); j != nil; j = j.Next() {
			if reflect.ValueOf(i.Value).Kind() == reflect.Int && reflect.ValueOf(j.Value).Kind() == reflect.Int {
				if j.Value.(int) < min.Value.(int) {
					min = j
				}
				continue
			}
			strI := ""
			if reflect.ValueOf(i.Value).Kind() == reflect.Int {
				strI = strconv.Itoa(i.Value.(int))
			} else {
				strI = i.Value.(string)
			}
			strM := ""
			if reflect.ValueOf(min.Value).Kind() == reflect.Int {
				strM = strconv.Itoa(min.Value.(int))
			} else {
				strM = min.Value.(string)
			}
			if strI < strM {
				min = j
			}
		}
		if min != i {
			i.Value, min.Value = min.Value, i.Value
		}
		bar.Add(1)
	}

	output, err := os.Create(filepath.Join("./sorted_data/selectionsort_" + file.Name()))
	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	writer := bufio.NewWriter(output)
	for i := data.Front(); i != nil; i = i.Next() {
		if reflect.ValueOf(i.Value).Kind() == reflect.Int {
			_, err := writer.WriteString(strconv.Itoa(i.Value.(int)) + "\n")
			if err != nil {
				fmt.Println("Error writing to file")
				return
			}
			continue
		}
		_, err := writer.WriteString(i.Value.(string) + "\n")
		if err != nil {
			fmt.Println("Error writing to file")
			return
		}
	}
	writer.Flush()
	output.Close()
	bar.Close()

}
