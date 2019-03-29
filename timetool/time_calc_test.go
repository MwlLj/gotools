package timetool

import (
	"fmt"
	"testing"
	"time"
)

func TestDaySub(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	// just one second
	t2, _ := time.Parse(layout, "2007-01-02 23:59:59")
	t1, _ := time.Parse(layout, "2007-01-03 00:00:00")
	fmt.Println(DaySub(t2, t1))
}

func TestDaySubByTime(t *testing.T) {
	fmt.Println(DaySubByTime("2007-01-02 23:59:59", "2007-01-03 00:00:00"))
}
