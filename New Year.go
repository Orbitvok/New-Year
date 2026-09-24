package main

import (
	"fmt"
	"time"
)

func newYear(date string) int {

	parse := "02.01.2006"

	dateParse, err := time.Parse(parse, date)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	nextYear := dateParse.Year() + 1

	newYearTime := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, dateParse.Location())
	days := int(newYearTime.Sub(dateParse).Hours()/24 - 1)

	return days
}

func main() {
	var date string

	_, err := fmt.Scan(&date)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("До New year %d дней", newYear(date))
}
