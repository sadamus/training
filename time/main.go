package main

import (
	"fmt"
	"time"
)

func diffDays(s string) int {
	now := time.Now()
	format := "2006-01-02 15:04:05"
	//fmt.Println(now)
	then, _ := time.Parse(format, fmt.Sprintf("%s %s", s, "22:00:00"))
	diff := now.Sub(then)
	//fmt.Println(then)
	return int(diff.Hours() / 24)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
func main() {
	// date is in format mm/dd/yyyy. Need it in yyyy-mm-dd format
	dt := "2018-06-15"
	println(diffDays(dt))

	then, _ := time.Parse("2006-01-02 00:00:00", "2018-06-15")
	fmt.Println(daysAgo(then))

}
