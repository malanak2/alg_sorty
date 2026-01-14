package sorts

import (
	"io/fs"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func HeapSort(file fs.DirEntry, doBar bool) time.Duration {
	data, err := util.LoadFileArr(file)
	if err != nil {
		panic(err)
	}
	time_s := time.Now()
	var bar *progressbar.ProgressBar
	if doBar {
		bar = progressbar.Default(int64(len(data)))
	}
	for i := len(data)/2 - 1; i >= 0; i-- {
		heapify(&data, len(data), i)
	}
	for i := len(data) - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		heapify(&data, i, 0)
		if doBar {
			bar.Add(1)
		}
	}
	if doBar {
		bar.Close()
	}
	time_e := time.Now()
	util.SaveResultArr(&data, file, "heapsort")
	return time_e.Sub(time_s)
}

func heapify(arr *[]any, heap_size int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < heap_size && util.Compare((*arr)[left], (*arr)[largest]) > 0 {
		largest = left
	}

	if right < heap_size && util.Compare((*arr)[right], (*arr)[largest]) > 0 {
		largest = right
	}

	if largest != i {
		(*arr)[i], (*arr)[largest] = (*arr)[largest], (*arr)[i]

		heapify(arr, heap_size, largest)
	}
}
