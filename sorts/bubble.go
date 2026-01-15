package sorts

import (
	"fmt"
	"io/fs"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func BubbleSort(file fs.DirEntry, doBar bool) time.Duration {
	data, err := util.LoadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sorting...")
	time_s := time.Now()
	var bar *progressbar.ProgressBar
	if doBar {
		bar = progressbar.Default(int64(data.Len()))
	}
	ran := true

	for ran {
		ran = false
		for i := data.Front(); i != nil && i.Next() != nil; i = i.Next() {
			if 0 < util.Compare(i.Value, i.Next().Value) {
				i.Value, i.Next().Value = i.Next().Value, i.Value
				ran = true
			}
		}
		if doBar {
			bar.Add(1)
		}
	}
	if doBar {
		bar.Close()
	}
	time_e := time.Now()
	util.SaveResult(data, file, "bubblesort")
	return time_e.Sub(time_s)
}
