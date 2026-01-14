package sorts

import (
	"io/fs"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func merge(data *[]any, left int, mid int, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	leftArr := make([]any, n1)
	rightArr := make([]any, n2)

	for i := range n1 {
		leftArr[i] = (*data)[left+i]
	}
	for i := range n2 {
		rightArr[i] = (*data)[mid+1+i]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if 0 >= util.Compare(leftArr[i], rightArr[j]) {
			(*data)[k] = leftArr[i]
			i += 1
		} else {
			(*data)[k] = rightArr[j]
			j += 1
		}
		k += 1
	}
	for i < n1 {
		(*data)[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		(*data)[k] = rightArr[j]
		j++
		k++
	}
}

func mergeSort(data *[]any, left int, right int, bar *progressbar.ProgressBar, doBar bool) {
	if left < right {
		mid := left + (right-left)/2

		mergeSort(data, left, mid, bar, doBar)
		mergeSort(data, mid+1, right, bar, doBar)
		merge(data, left, mid, right)
		if doBar {
			bar.Add(1)
		}
	}
}

func MergeSort(file fs.DirEntry, doBar bool) time.Duration {
	data, err := util.LoadFileArr(file)
	if err != nil {
		panic(err)
	}
	time_s := time.Now()
	var bar *progressbar.ProgressBar
	if doBar {
		bar = progressbar.Default(int64(len(data)))
	}
	mergeSort(&data, 0, len(data)-1, bar, doBar)
	if doBar {
		bar.Close()
	}
	time_e := time.Now()
	util.SaveResultArr(&data, file, "mergeSort")
	return time_e.Sub(time_s)
}
