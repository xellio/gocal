package gocal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Cal struct {
	FromDate       time.Time
	ToDate         time.Time
	FirstDayOfWeek int
	MarkToday      bool
	HideHeader     bool
	ColorDefault   string
	ColorToday     string
	ColorHighlight string
	Marker         []time.Time
}

func (c *Cal) init() {

	if c.FromDate.IsZero() {
		c.FromDate = time.Now()
	}

	if c.ToDate.IsZero() {
		c.ToDate = c.FromDate
	}

	if c.ColorToday == "" {
		c.ColorToday = COLOR_TODAY
	}

	if c.ColorDefault == "" {
		c.ColorDefault = COLOR_DEFAULT
	}

	if c.ColorHighlight == "" {
		c.ColorHighlight = COLOR_HIGHLIGHT
	}
}

func (c *Cal) Print() error {
	c.init()
	if !c.HideHeader {
		fmt.Println("---------------------------")
		fmt.Println(c.FromDate.Day(), c.FromDate.Month(), c.FromDate.Year())
		fmt.Println("---------------------------")
	}

	dateYear, dateMonth, _ := c.FromDate.Date()
	dateLocation := c.FromDate.Location()
	firstOfMonth := time.Date(dateYear, dateMonth, 1, 0, 0, 0, 0, dateLocation)

	weeks, err := c.calculateWeeks(firstOfMonth)
	if err != nil {
		return err
	}
	if err = c.printCalendar(weeks); err != nil {
		return err
	}

	return nil
}

func (c *Cal) calculateWeeks(firstOfMonth time.Time) ([][]int, error) {
	var weeks [][]int

	var days []int
	for i := 0; i < int(firstOfMonth.Weekday())-c.FirstDayOfWeek; i++ {
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

func (c *Cal) printCalendar(weeks [][]int) error {

	if err := c.printWeekdayHeader(); err != nil {
		return err
	}
	c.printWeeks(weeks)
	return nil
}

func (c *Cal) printWeekdayHeader() error {

	var orderedWeekDays []string
	orderedWeekDays = DAYS[c.FirstDayOfWeek:len(DAYS)]

	if c.FirstDayOfWeek != 0 {
		orderedWeekDays = append(orderedWeekDays, DAYS[0:c.FirstDayOfWeek]...)
	}

	fmt.Println(strings.Join(orderedWeekDays, " "))
	return nil
}

func (c *Cal) printWeeks(weeks [][]int) error {
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
				fmt.Printf("\033["+c.ColorToday+"m%s \033[0m", dayToPrint)
			} else {
				fmt.Printf("\033["+c.ColorDefault+"m%s \033[0m", dayToPrint)
			}

		}
		fmt.Print("\n")
	}
	return nil
}
