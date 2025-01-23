package utils

import (
	"fmt"
	"time"
)

func isLeapYear(date time.Time) bool {
	return date.Year()%4 == 0
}

// CalcDays returns number of days elapsed between two dates
// Used to calculate swaps that are applied daily
// Accepts two params: start, and end time.Time
// Does not return negative days
func CalcDays(start, end time.Time) uint16 {

	var days uint16 = 0

	if start.Year() == end.Year() {
		return uint16(end.YearDay()) - uint16(start.YearDay())
	} else {

		var years []int

		for i := start.Year(); i <= end.Year(); i++ {
			years = append(years, i)
		}

		for _, v := range years {

			if start.Year() == v {
				fmt.Println(time.Date(v, 1, 1, 1, 0, 0, 0, time.UTC), "vv")
				if isLeapYear(time.Date(v, 1, 1, 1, 0, 0, 0, time.UTC)) {
					days += 366 - uint16(start.YearDay())
					fmt.Println("leap")
				} else {
					days += 365 - uint16(start.YearDay())
				}

			}

			if v != start.Year() && v != end.Year() {
				fmt.Println(time.Date(v, 1, 1, 1, 0, 0, 0, time.UTC), "vvv")
				if isLeapYear(time.Date(v, 1, 1, 1, 0, 0, 0, time.UTC)) {
					fmt.Println("leap mid")
					days += 366
				} else {
					days += 365
				}

			}

			if end.Year() == v {
				fmt.Println(v, "vv1")
				days += uint16(end.YearDay())
			}

			fmt.Println(v)
		}
	}

	return days
}
