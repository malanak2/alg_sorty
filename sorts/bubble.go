package sorts

import (
	"fmt"
	"io/fs"
	"reflect"
	"strconv"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func BubbleSort(file fs.DirEntry) {
	data, err := util.LoadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sorting... (bar shows worst case scenario of operations)")

	bar := progressbar.Default(int64(data.Len()) * int64(data.Len()))
	ran := true

	for ran {
		ran = false
		for i := data.Front(); i != nil && i.Next() != nil; i = i.Next() {
			switch v := i.Value.(type) {
			case int:
				if i.Value.(int) > i.Next().Value.(int) {
					i.Value, i.Next().Value = i.Next().Value, i.Value
					ran = true
				}
				continue
			case string:
				strI := ""
				switch w := i.Value.(type) {
				case int:
					strI = strconv.Itoa(i.Value.(int))
				case string:
					strI = i.Value.(string)
				default:
					fmt.Println("Unknown type:", reflect.TypeOf(w))
					return
				}

				strNext := ""
				switch w := i.Next().Value.(type) {
				case string:
					strNext = strconv.Itoa(i.Next().Value.(int))
				case int:
					strNext = i.Next().Value.(string)
				default:
					fmt.Println("Unknown type:", reflect.TypeOf(w))
					return
				}

				if strI > strNext {
					i.Value, i.Next().Value = i.Next().Value, i.Value
					ran = true
				}

				bar.Add(1)
			default:
				fmt.Println("Unknown type:", reflect.TypeOf(v))
				return
			}
		}
	}
	bar.Close()
	util.SaveResult(data, file, "bubblesort")

}
