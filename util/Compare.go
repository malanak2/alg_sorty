package util

import (
	"strconv"
)

func Compare(a any, b any) int {
	strA := ""
	strB := ""
	switch va := a.(type) {
	case float64:
		switch vb := b.(type) {
		case float64:
			if va == vb {
				return 0
			} else if va < vb {
				return -1
			}
			return 1
		case string:
			strA = strconv.FormatFloat(va, 'f', -1, 64)
			strB = vb
		}
	case string:
		strA = va
		switch vb := b.(type) {
		case float64:
			strB = strconv.FormatFloat(vb, 'f', -1, 64)
		case string:
			if a.(string) == vb {
				return 0
			} else if a.(string) < vb {
				return -1
			}
			return 1
		}
	}
	if strA == strB {
		return 0
	} else if strA < strB {
		return -1
	}
	return 1
}
