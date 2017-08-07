/**
 * package gocal
 */
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
	monthsToPrint  []time.Time
}

// initialize the Cal struct
// set default values and collects/calculates data for printing
func (c *Cal) init() {

	if c.FromDate.IsZero() {
		c.FromDate = time.Now()
	}
	c.FromDate = time.Date(c.FromDate.Year(), c.FromDate.Month(), 1, 0, 0, 0, 0, c.FromDate.Location())

	if c.ToDate.IsZero() {
		//@todo: parse Marker slice for ToDate value
		c.ToDate = c.FromDate.AddDate(0, 1, -1)
	} else {
		c.ToDate = time.Date(c.ToDate.Year(), c.ToDate.Month()+1, 0, 0, 0, 0, 0, c.ToDate.Location())
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

	for d := c.FromDate; d.Month() <= c.ToDate.Month(); d = d.AddDate(0, 1, 0) {
		c.monthsToPrint = append(c.monthsToPrint, d)
	}

}

// Print the calendar
func (c *Cal) Print() error {
	c.init()

	for _, month := range c.monthsToPrint {
		if !c.HideHeader {
			fmt.Println("---------------------------")
			fmt.Println(month.Month(), month.Year())
			fmt.Println("---------------------------")
		}

		weeks, err := c.calculateWeeks(month)
		if err != nil {
			return err
		}
		if err = c.printCalendar(weeks); err != nil {
			return err
		}
	}

	return nil
}

// Returns a slice containing another slice with all time.Time values for one week (line)
func (c *Cal) calculateWeeks(firstOfMonth time.Time) ([][]time.Time, error) {
	var weeks [][]time.Time

	var days []time.Time

	slotsToFill := int(firstOfMonth.Weekday()) - c.FirstDayOfWeek
	// if the 1st is a sunday and week starts on monday, we have to fill 6 slots
	// slotsToFill (in this case) is -1
	if slotsToFill < 0 {
		slotsToFill = slotsToFill + 7
	}
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

// Triggers the printing of the calendar header and the weeks
func (c *Cal) printCalendar(weeks [][]time.Time) error {

	if err := c.printWeekdayHeader(); err != nil {
		return err
	}
	c.printWeeks(weeks)
	return nil
}

// Print the calendar header
func (c *Cal) printWeekdayHeader() error {

	var orderedWeekDays []string
	orderedWeekDays = DAYS[c.FirstDayOfWeek:len(DAYS)]

	if c.FirstDayOfWeek != 0 {
		orderedWeekDays = append(orderedWeekDays, DAYS[0:c.FirstDayOfWeek]...)
	}

	fmt.Println(strings.Join(orderedWeekDays, " "))
	return nil
}

// Print the weeks
func (c *Cal) printWeeks(weeks [][]time.Time) error {
	today := time.Now()

	for _, days := range weeks {
		for _, day := range days {
			printFormat := "\033[" + c.ColorDefault + "m%s \033[0m"
			dayToPrint := " " + strconv.Itoa(day.Day())

			if day.Day() < 10 {
				dayToPrint = " " + dayToPrint
			}

			if len(c.Marker) > 0 && c.shouldBeMarked(day) {
				printFormat = "\033[" + c.ColorHighlight + "m%s \033[0m"
			}

			if c.MarkToday && today.Day() == day.Day() &&
				today.Month() == day.Month() &&
				today.Year() == day.Year() {
				printFormat = "\033[" + c.ColorToday + "m%s \033[0m"
			}

			fmt.Printf(printFormat, dayToPrint)

		}
		fmt.Print("\n")
	}
	return nil
}

// checks if the given time.Time is in the c.Marker slice
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
