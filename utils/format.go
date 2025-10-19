package utils

import "time"

func ToDateTime(timeValue *time.Time) string {
	return timeValue.Format("2006-01-02 15:04:05")
}
