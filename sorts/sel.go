package sorts

import (
	"io/fs"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func SelectionSort(file fs.DirEntry, doBar bool) time.Duration {
	data, err := util.LoadFile(file)
	if err != nil {
		panic(err)
	}
	time_s := time.Now()
	var bar *progressbar.ProgressBar
	if doBar {

		bar = progressbar.Default(int64(data.Len()))
	}
	for i := data.Front(); i != nil; i = i.Next() {
		min := i
		for j := i.Next(); j != nil; j = j.Next() {
			if 0 < util.Compare(i.Value, j.Value) {
				min = j
			}
			continue
		}
		if min != i {
			i.Value, min.Value = min.Value, i.Value
		}
		if doBar {

			bar.Add(1)
		}
	}

	if doBar {
		bar.Close()
	}
	time_e := time.Now()

	util.SaveResult(data, file, "selectionSort")
	return time_e.Sub(time_s)
}
