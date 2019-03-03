package timetool

import (
	"fmt"
	"testing"
)

func TestGetNowDayFormat(t *testing.T) {
	now := GetNowDayFormat()
	fmt.Println(now)
}

func TestGetNowSecondFormat(t *testing.T) {
	now := GetNowSecondFormat()
	fmt.Println(now)
}
