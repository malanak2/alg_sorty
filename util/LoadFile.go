package util

import (
	"bufio"
	"container/list"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

func LoadFile(file fs.DirEntry) (*list.List, error) {
	data := list.New()
	open, err := os.Open(filepath.Join("./data", file.Name()))
	if err == nil {
		defer open.Close()
	} else {
		fmt.Println("Error opening file")
		return list.New(), err
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
		return list.New(), err
	}
	return data, nil
}
