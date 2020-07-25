package date_utils

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	now := GetNow()
	return now.Format(apiDateLayout)
}

func GetNowDbFormat() string {
	now := GetNow()
	return now.Format(apiDbLayout)
}
