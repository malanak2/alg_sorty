package sorts

import (
	"io/fs"
	"math"
	"strconv"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func findMax(data *[]int) int {
	max := 0
	largest := 0

	for i := 1; i < len(*data); i++ {
		ln := len(strconv.Itoa((*data)[i]))
		if ln > max {
			max = ln
			largest = (*data)[i]
		}
	}
	return largest
}

func RadixSort(file fs.DirEntry, doBar bool) time.Duration {
	dataA, err := util.LoadFileArr(file)
	if err != nil {
		panic(err)
	}

	data := make([]int, len(dataA))
	for i := range dataA {
		data[i] = int(math.Round(dataA[i].(float64)))
	}

	time_s := time.Now()
	var bar *progressbar.ProgressBar
	if doBar {
		bar = progressbar.Default(int64(len(data)))
	}
	maxI := findMax(&data)
	exp := 1
	for maxI/exp >= 1 {
		n := len(data)
		output := make([]int, n)
		count := make([]int, 10)

		for i := 0; i < n; i++ {
			index := (data[i] / exp) % 10
			count[index]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			index := (data[i] / exp) % 10
			output[count[index]-1] = data[i]
			if doBar {
				bar.Add(1)
			}
			count[index]--
		}

		copy(data, output)
		exp *= 10
	}
	if doBar {
		bar.Close()
	}
	time_e := time.Now()
	dataSave := make([]any, len(data))
	for i := range data {
		dataSave[i] = float64(data[i])
	}
	util.SaveResultArr(&dataSave, file, "radixsort")
	return time_e.Sub(time_s)
}
