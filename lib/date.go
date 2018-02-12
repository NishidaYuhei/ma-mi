package lib

import "time"

func GetNowYearAndMonth() string {
	return time.Now().Format("200601")
}

func GetDateStr() string {
	return time.Now().Format("200601021504")
}
