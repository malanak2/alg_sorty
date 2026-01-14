package sorts

import (
	"io/fs"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func partition(arr *[]any, low int, high int) int {
	pivot := (*arr)[high]
	i := low - 1

	for j := low; j < high; j++ {
		if util.Compare((*arr)[j], pivot) < 0 {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]
	return i + 1
}

func quicksort(arr *[]any, low int, high int, bar *progressbar.ProgressBar, doBar bool) {
	if low < high {
		pi := partition(arr, low, high)

		// Wordss...
		if doBar {
			bar.Add(1)
		}
		quicksort(arr, low, pi-1, bar, doBar)
		quicksort(arr, pi+1, high, bar, doBar)
	}
}

func QuickSort(file fs.DirEntry, doBar bool) time.Duration {
	data, err := util.LoadFileArr(file)
	if err != nil {
		panic(err)
	}
	time_s := time.Now()
	var bar *progressbar.ProgressBar
	if doBar {
		bar = progressbar.Default(int64(len(data)))
	}

	quicksort(&data, 0, len(data)-1, bar, doBar)

	if doBar {
		bar.Close()
	}
	time_e := time.Now()
	util.SaveResultArr(&data, file, "quicksort")
	return time_e.Sub(time_s)
}
