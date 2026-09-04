package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	nextYear := now.Year() + 1

	newYearTime := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, now.Location())

	days := int(newYearTime.Sub(now).Hours() / 24)

	fmt.Printf("До New year %d день", days)
}
