package sorts

import (
	"io/fs"
	"strconv"
	"time"

	"github.com/malanak2/alg_sorts/util"
	"github.com/schollz/progressbar/v3"
)

func findMaxLen(data []string) int {
	max := 0
	for _, s := range data {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

func RadixStringSort(file fs.DirEntry, doBar bool) time.Duration {
	dataA, err := util.LoadFileArr(file)
	if err != nil {
		panic(err)
	}

	data := make([]string, len(dataA))
	for i := range dataA {
		switch s := dataA[i].(type) {
		case float64:
			data[i] = strconv.FormatFloat(s, 'f', -1, 64)
		case string:
			data[i] = s
		}
	}

	time_s := time.Now()
	maxLen := findMaxLen(data)
	n := len(data)

	var bar *progressbar.ProgressBar
	if doBar {
		bar = progressbar.Default(int64(maxLen * n))
	}

	for pos := maxLen - 1; pos >= 0; pos-- {
		output := make([]string, n)
		count := make([]int, 256) // ASCII

		for i := range n {
			charVal := 0
			if pos < len(data[i]) {
				charVal = int(data[i][pos])
			}
			count[charVal]++
		}

		for i := 1; i < 256; i++ {
			count[i] += count[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			charVal := 0
			if pos < len(data[i]) {
				charVal = int(data[i][pos])
			}
			output[count[charVal]-1] = data[i]
			count[charVal]--

			if doBar {
				bar.Add(1)
			}
		}

		copy(data, output)
	}

	if doBar {
		bar.Close()
	}
	time_e := time.Now()

	dataSave := make([]any, len(data))
	for i := range data {
		dataSave[i] = data[i]
	}
	util.SaveResultArr(&dataSave, file, "radixstrsort")

	return time_e.Sub(time_s)
}

