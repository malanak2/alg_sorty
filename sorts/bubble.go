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

	"github.com/schollz/progressbar/v3"
)

func BubbleSort(file fs.DirEntry) {
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

	fmt.Println("Sorting... (bar shows worst case scenario of operations)")

	bar := progressbar.Default(int64(data.Len()) * int64(data.Len()))
	ran := true

	for ran {
		ran = false
		for i := data.Front(); i != nil && i.Next() != nil; i = i.Next() {
			if reflect.ValueOf(i.Value).Kind() == reflect.Int && reflect.ValueOf(i.Next().Value).Kind() == reflect.Int {
				if i.Value.(int) > i.Next().Value.(int) {
					i.Value, i.Next().Value = i.Next().Value, i.Value
					ran = true
				}
				continue
			}
			strI := ""
			if reflect.ValueOf(i.Value).Kind() == reflect.Int {
				strI = strconv.Itoa(i.Value.(int))
			} else {
				strI = i.Value.(string)
			}
			strNext := ""
			if reflect.ValueOf(i.Next().Value).Kind() == reflect.Int {
				strNext = strconv.Itoa(i.Next().Value.(int))
			} else {
				strNext = i.Next().Value.(string)
			}
			if strI > strNext {
				i.Value, i.Next().Value = i.Next().Value, i.Value
				ran = true
			}
		}
		bar.Add(1)
	}

	output, err := os.Create(filepath.Join("./sorted_data/bubblesort_" + file.Name()))
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
