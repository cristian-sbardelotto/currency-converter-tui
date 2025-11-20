package utils

import (
	"math"
	"strconv"
)

func FormatFloat(num float64) string {
	rounded := math.Round(num*100) / 100
	return strconv.FormatFloat(rounded, 'f', -1, 64)
}
