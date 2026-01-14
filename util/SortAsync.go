package util

import (
	"io/fs"
	"time"
)

type sort func(fs.DirEntry, bool) time.Duration

func SortAsync(fun sort, file fs.DirEntry, doBar bool, ch chan time.Duration) {
	duration := fun(file, doBar)
	ch <- duration
}
