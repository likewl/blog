package middleware

import (
	"time"
)

func GetTime() string {
	time := time.Now().Format("2006-01-02 15:04:05")
	return time
}
func GetTimes() (int, int, int) {
	date, month, day := time.Now().Date()
	return date, int(month), day
}
