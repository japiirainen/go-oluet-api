package helpers

import (
	"strconv"
	"time"
)

//ToFloat converts string values to float64
func ToFloat(v string) float64 {
	res, _ := strconv.ParseFloat(v, 64)
	return res
}

//ToInt converts string values to int
func ToInt(v string) int {
	res, _ := strconv.Atoi(v)
	return res
}

//ParseTime converts string values to time.Time
func ParseTime(timeStr string) time.Time {
	time, _ := time.Parse("2006-Jan-02", timeStr)
	return time
}
