package sorts

import (
	"fmt"
	"io/fs"
	"reflect"
	"strconv"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func InsertionSort(file fs.DirEntry) {
	data, err := util.LoadFile(file)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sorting... (bar shows worst case scenario of operations)")

	bar := progressbar.Default(int64(data.Len()))

	curr := data.Front().Next()
	next := curr.Next()
	// while
	for curr != nil {
		if curr.Prev() == nil {
			curr = next
			if curr.Next() == nil {
				break
			}
			next = curr.Next()
			continue
		}
		switch v := curr.Value.(type) {
		case int:
			if curr.Prev().Value.(int) < curr.Value.(int) {
				curr = next
				if curr == nil {
					break
				}
				next = curr.Next()
				bar.Add(1)
				continue
			} else {
				data.MoveAfter(curr.Prev(), curr)
				continue
			}
		case string:
			strCurr := ""
			if reflect.ValueOf(curr.Value).Kind() == reflect.Int {
				strCurr = strconv.Itoa(curr.Value.(int))
			} else {
				strCurr = curr.Value.(string)
			}
			strPrev := ""
			if reflect.ValueOf(curr.Prev().Value).Kind() == reflect.Int {
				strPrev = strconv.Itoa(curr.Prev().Value.(int))
			} else {
				strPrev = curr.Prev().Value.(string)
			}
			if strPrev < strCurr {
				curr = next
				if curr == nil {
					break
				}
				next = curr.Next()
				bar.Add(1)
				continue
			} else {
				data.MoveAfter(curr.Prev(), curr)
				continue
			}
		default:
			fmt.Println("Unknown type:", reflect.TypeOf(v))
			return
		}

	}

	bar.Close()
	util.SaveResult(data, file, "insertionsort")
}
