package sorts

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func SelectionSort(file fs.DirEntry) {
	data, err := util.LoadFile(file)
	if err != nil {
		panic(err)
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

	bar.Close()

	util.SaveResult(data, file, "selectionSort")
}
