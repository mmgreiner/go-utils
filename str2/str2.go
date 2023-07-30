package str2

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

/*
in importing module, for local module import, add:

go.mod:
replace mmgreiner/str2 => ../str2
*/

var Warn = false

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func logCannotParse(err error, s string, pos any) {
	if Warn && err != nil {
		msg := fmt.Sprint(err)
		if pos != nil {
			msg = msg + fmt.Sprint(" at ", pos)
		}
		log.Println(msg)
	}
}
func TraceToInt(s string, pos any) int {
	v, err := strconv.ParseInt(s, 10, 32)
	logCannotParse(err, s, pos)
	return int(v)
}

func ToInt(s string) int {
	return TraceToInt(s, nil)
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func TraceToFloat(s string, pos any) float64 {
	v, err := strconv.ParseFloat(s, 64)
	logCannotParse(err, s, pos)
	return v
}

func ToFloat(s string) float64 {
	return TraceToFloat(s, -1)
}

func IsBool(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

func TraceToBool(s string, pos any) bool {
	v, err := strconv.ParseBool(s)
	logCannotParse(err, s, pos)
	return v
}

func ToBool(s string) bool {
	return TraceToBool(s, -1)
}

// https://pkg.go.dev/time#pkg-constants
var formats = []string{
	"01.02.2006",
	"01.02.06",
	"01-02-06",
	"02.01.2006 15:04:05", // 13.01.2023  00:00:00
	time.DateOnly,
	time.RFC3339,
}

func toTime(s string) (time.Time, error) {
	for _, f := range formats {
		d, err := time.Parse(f, s)
		if err == nil {
			return d, nil
		}
	}
	return time.Time{}, errors.New("cannot parse date " + s)
}

func IsTime(s string) bool {
	_, err := toTime(s)
	return err == nil
}

func TraceToTime(s string, pos any) time.Time {
	v, err := toTime(s)
	logCannotParse(err, s, pos)
	return v
}

func ToTime(s string) time.Time {
	return TraceToTime(s, -1)
}
