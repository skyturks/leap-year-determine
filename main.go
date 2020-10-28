package main

import "fmt"

func main() {
	year := 2021

	fmt.Println(year, "is leap year? ", isLeapYear(year))
}

func isLeapYear(year int) bool {
	isLeap := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				isLeap = true
			} else {
				isLeap = false
			}
		} else {
			isLeap = true
		}
	} else {
		isLeap = false
	}
	return isLeap
}
