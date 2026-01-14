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
		if reflect.ValueOf(i.Value).Kind() == reflect.Float64 {
			_, err := writer.WriteString(strconv.FormatFloat(i.Value.(float64), 'f', -1, 64) + "\n")
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

func SaveResultArr(data *[]any, file fs.DirEntry, sortName string) error {
	output, err := os.Create(filepath.Join("./sorted_data/" + sortName + "_" + file.Name()))
	if err != nil {
		fmt.Println("Error creating file")
		return err
	}

	writer := bufio.NewWriter(output)
	for i := 0; i != len(*data); i++ {
		switch (*data)[i].(type) {
		case float64:

			_, err := writer.WriteString(strconv.FormatFloat((*data)[i].(float64), 'f', -1, 64) + "\n")
			if err != nil {
				fmt.Println("Error writing to file")
				return err
			}
			continue
		case string:
			_, err := writer.WriteString((*data)[i].(string) + "\n")
			if err != nil {
				fmt.Println("Error writing to file")
				return err
			}
		}
	}
	writer.Flush()
	output.Close()
	return nil
}
