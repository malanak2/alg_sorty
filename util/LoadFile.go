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
		lineI, err := strconv.ParseFloat(line, 64)
		if err == nil {
			data.PushBack(lineI)
			continue
		}
		data.PushBack(line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s", err)
		return list.New(), err
	}
	return data, nil
}

func LoadFileArr(file fs.DirEntry) ([]any, error) {
	data := []any{}
	open, err := os.Open(filepath.Join("./data", file.Name()))
	if err == nil {
		defer open.Close()
	} else {
		fmt.Println("Error opening file")
		return []any{}, err
	}
	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(open)

	// Loop through the file and read each line
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		lineI, err := strconv.ParseFloat(line, 64)
		if err == nil {
			data = append(data, lineI)
			continue
		}
		data = append(data, line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s", err)
		return []any{}, err
	}
	return data, nil
}
