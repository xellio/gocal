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

//
// Cal struct stores settings for the calendar
//
type Cal struct {
	FromDate       time.Time   // Specify the month to display. If no value is passed. time.Now() will be used. NOTE: gocal will use the month - not the day. Passing 15th of August will be the same as 1st of August.
	ToDate         time.Time   // Specify the last month to display. If no value is passed, only the month of Cal.FromDate is printed.
	FirstDayOfWeek int         // By default, a week starts with Sunday. If you want to start with Monday (or any other day), you can set this value to 1 (for Monday).
	MarkToday      bool        // Setting this flag to true, the todays date is highlighted in the Cal.ColorToday color.
	HideHeader     bool        // Setting this flag to false will hide the output-header and display only the calendar without any other information.
	ColorDefault   string      // Default: 29 Specify the default output color ([ANSI Color Codes](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors))
	ColorToday     string      // Default: 31 Specify the default output color ([ANSI Color Codes](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors))
	ColorHighlight string      // Default: 32 Specify the default output color ([ANSI Color Codes](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors))
	Marker         []time.Time // []time.Time slice of dates to highlight in the calendar.
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
		c.ColorToday = ColorToday
	}

	if c.ColorDefault == "" {
		c.ColorDefault = ColorDefault
	}

	if c.ColorHighlight == "" {
		c.ColorHighlight = ColorHighlight
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

// Print the calendar using c attributes
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
	orderedWeekDays = Days[c.FirstDayOfWeek:len(Days)]

	if c.FirstDayOfWeek != 0 {
		orderedWeekDays = append(orderedWeekDays, Days[0:c.FirstDayOfWeek]...)
	}

	fmt.Println(strings.Join(orderedWeekDays, " "))
	return nil
}

// Print the weeks
func (c *Cal) printWeeks(weeks [][]time.Time) error {
	today := time.Now()

	for wc, days := range weeks {
		for _, day := range days {
			printFormat := "\033[" + c.ColorDefault + "m%s \033[0m"
			dayToPrint := " " + strconv.Itoa(day.Day())

			// first week-line, but day from previous month - print spaces instead of value
			if wc == 0 && day.Day() > 7 {
				dayToPrint = "   "
			}

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
