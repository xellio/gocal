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
	c.FromDate = time.Date(c.FromDate.Year(), c.FromDate.Month(), 1, 0, 0, 0, 0, c.FromDate.Location())

	if c.ToDate.IsZero() {
		c.ToDate = c.FromDate.AddDate(0, 1, 0)
	}
	c.ToDate = time.Date(c.ToDate.Year(), c.ToDate.Month()+1, 0, 0, 0, 0, 0, c.ToDate.Location())

	if c.ColorToday == "" {
		c.ColorToday = COLOR_TODAY
	}

	if c.ColorDefault == "" {
		c.ColorDefault = COLOR_DEFAULT
	}

	if c.ColorHighlight == "" {
		c.ColorHighlight = COLOR_HIGHLIGHT
	}

	if len(c.Marker) > 0 {
		// filter marker in timeframe
		var filteredMarker []time.Time
		for _, d := range c.Marker {
			if d.After(c.FromDate) && d.Before(c.ToDate) {
				filteredMarker = append(filteredMarker, d)
			}
		}
		c.Marker = filteredMarker
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

func (c *Cal) calculateWeeks(firstOfMonth time.Time) ([][]time.Time, error) {
	var weeks [][]time.Time

	var days []time.Time

	slotsToFill := int(firstOfMonth.Weekday()) - c.FirstDayOfWeek
	for i := slotsToFill; i > 0; i-- {
		dateToAppend := firstOfMonth.AddDate(0, 0, -i)
		days = append(days, dateToAppend)
	}

	for d := firstOfMonth; d.Month() == firstOfMonth.Month(); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
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

func (c *Cal) printCalendar(weeks [][]time.Time) error {

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

func (c *Cal) printWeeks(weeks [][]time.Time) error {
	today := time.Now()
	for _, days := range weeks {
		for _, day := range days {
			printFormat := "\033[" + c.ColorDefault + "m%s \033[0m"
			dayToPrint := " " + strconv.Itoa(day.Day())

			if day.Month() != c.FromDate.Month() {
				dayToPrint = "   "
			}
			if day.Day() < 10 {
				dayToPrint = " " + dayToPrint
			}

			if len(c.Marker) > 0 && c.shouldBeMarked(day) {
				printFormat = "\033[" + c.ColorHighlight + "m%s \033[0m"
			}

			if today.Day() == day.Day() {
				printFormat = "\033[" + c.ColorToday + "m%s \033[0m"
			}

			fmt.Printf(printFormat, dayToPrint)

		}
		fmt.Print("\n")
	}
	return nil
}

func (c *Cal) shouldBeMarked(day time.Time) bool {
	for _, marker := range c.Marker {
		if marker.Day() == day.Day() &&
			marker.Month() == day.Month() &&
			marker.Year() == day.Year() {
			return true
		}
	}
	return false
}
