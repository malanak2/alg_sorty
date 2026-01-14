package sorts

import (
	"fmt"
	"io/fs"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func InsertionSort(file fs.DirEntry, doBar bool) time.Duration {
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
		if 0 > util.Compare(curr.Prev().Value, curr.Value) {
			curr = next
			if curr == nil {
				break
			}
			next = curr.Next()
			if doBar {
				bar.Add(1)
			}
			continue
		} else {
			data.MoveAfter(curr.Prev(), curr)
			continue
		}
	}

	if doBar {
		bar.Close()
	}
	time_e := time.Now()
	util.SaveResult(data, file, "insertionsort")
	return time_e.Sub(time_s)
}
