package helpers

import (
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
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
	t := strings.Split(timeStr, ".")
	dd := t[0]
	mm := t[1]
	yyyy := t[2]
	formatted := yyyy + "-" + mm + "-" + dd
	myDate, _ := time.Parse("2006-01-02", formatted)
	return myDate
}

//Duration prints the time of some func
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	log.Infof("%s lasted %s", name, elapsed)
}
