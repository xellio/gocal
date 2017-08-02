package gocal

import (
	"fmt"
	"strconv"
	"time"
)

func Print(date time.Time) error {
	fmt.Println("--------------------------------")
	fmt.Println(date.Month(), date.Year())
	fmt.Println("--------------------------------")

	dateYear, dateMonth, _ := date.Date()
	dateLocation := date.Location()
	firstOfMonth := time.Date(dateYear, dateMonth, 1, 0, 0, 0, 0, dateLocation)

	weeks, err := calculateWeeks(firstOfMonth)
	if err != nil {
		return err
	}
	if err = printCalendar(weeks); err != nil {
		return err
	}

	return nil
}

func calculateWeeks(firstOfMonth time.Time) ([][]int, error) {
	var weeks [][]int

	var days []int
	for i := 0; i < int(firstOfMonth.Weekday()); i++ {
		days = append(days, 0)
	}

	for d := firstOfMonth; d.Month() == firstOfMonth.Month(); d = d.AddDate(0, 0, 1) {
		days = append(days, d.Day())
		if len(days) == 7 {
			weeks = append(weeks, days)
			days = nil
		}
	}

	if len(days) > 0 {
		weeks = append(weeks, days)
	}
	return weeks, nil
}

func printCalendar(weeks [][]int) error {

	if err := printWeekdayHeader(); err != nil {
		return err
	}
	printWeeks(weeks)
	return nil
}

func printWeekdayHeader() error {
	var weekdayHeader string
	for day, _ := range DAYS {
		weekdayHeader += day + " "
	}
	fmt.Println(weekdayHeader)
	return nil
}

func printWeeks(weeks [][]int) error {
	today := time.Now()
	for _, days := range weeks {
		for _, day := range days {
			dayToPrint := strconv.Itoa(day)
			if day == 0 {
				dayToPrint = " "
			}
			if day < 10 {
				dayToPrint = "  " + dayToPrint
			} else {
				dayToPrint = " " + dayToPrint
			}

			if today.Day() == day {
				fmt.Printf("\033["+COLOR_TODAY+"m%s \033[0m", dayToPrint)
			} else {
				fmt.Printf("\033["+COLOR_DEFAULT+"m%s \033[0m", dayToPrint)
			}

		}
		fmt.Print("\n")
	}
	return nil
}
