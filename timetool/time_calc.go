package timetool

import (
	"time"
)

func DaySub(t1, t2 time.Time) int {
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	return int(t1.Sub(t2).Hours() / 24)
}

func DaySubByTime(t1, t2 string) int {
	layout := "2006-01-02 15:04:05"
	time1, _ := time.Parse(layout, t1)
	time2, _ := time.Parse(layout, t2)
	return DaySub(time1, time2)
}
