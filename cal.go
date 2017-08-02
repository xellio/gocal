package gocal

import (
	"fmt"
	"strconv"
	"time"
)

func Print(date time.Time) error {
	fmt.Println("------------------------")
	fmt.Println(date.Month(), date.Year())
	fmt.Println("------------------------")

	dateYear, dateMonth, _ := date.Date()
	dateLocation := date.Location()

	firstOfMonth := time.Date(dateYear, dateMonth, 1, 0, 0, 0, 0, dateLocation)

	if err := printWeekdayHeader(); err != nil {
		return err
	}

	var lines []string
	for i := 0; i < int(firstOfMonth.Weekday()); i++ {
		lines = append(lines, " ")
	}

	fmt.Println(firstOfMonth.Weekday())
	for d := firstOfMonth; d.Month() == firstOfMonth.Month(); d = d.AddDate(0, 0, 1) {
		lines = append(lines, strconv.Itoa(d.Day()))
	}

	fmt.Println(lines)

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
