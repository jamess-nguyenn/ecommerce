package helpers

import "time"

func GetTimeNow() time.Time {
	return time.Now()
}

func GetDatetime() string {
	return GetTimeNow().Format(time.DateTime)
}

func GetDate() string {
	return GetTimeNow().Format(time.DateOnly)
}
