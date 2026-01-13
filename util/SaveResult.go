package util

import (
	"bufio"
	"container/list"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

func SaveResult(data *list.List, file fs.DirEntry, sortName string) error {
	output, err := os.Create(filepath.Join("./sorted_data/" + sortName + "_" + file.Name()))
	if err != nil {
		fmt.Println("Error creating file")
		return err
	}

	writer := bufio.NewWriter(output)
	for i := data.Front(); i != nil; i = i.Next() {
		if reflect.ValueOf(i.Value).Kind() == reflect.Int {
			_, err := writer.WriteString(strconv.Itoa(i.Value.(int)) + "\n")
			if err != nil {
				fmt.Println("Error writing to file")
				return err
			}
			continue
		}
		_, err := writer.WriteString(i.Value.(string) + "\n")
		if err != nil {
			fmt.Println("Error writing to file")
			return err
		}
	}
	writer.Flush()
	output.Close()
	return nil
}
