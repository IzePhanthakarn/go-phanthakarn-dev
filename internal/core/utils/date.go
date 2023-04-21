package utils

import "time"

func StringToDate(s string) (time.Time, error) {
	formatDate := "2006-01-02"
	return time.Parse(formatDate, s)
}

func StringToDateTime(s string) (time.Time, error) {
	formatDateTime := "2006-01-02 15:04:05"
	return time.Parse(formatDateTime, s)
}
