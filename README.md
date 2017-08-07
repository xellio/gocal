## gocal - a cli calendar presentation

## Motivation
Needed a calendar for my mk project - so here it is

## Usage
```
go get -u github.com/xellio/gocal
```
### Print calendar
```
calendar := gocal.Cal{}
calendar.Print()
```
### Result
```
---------------------------
August 2017
---------------------------
MON TUE WED THU FRI SAT SUN
      1   2   3   4   5   6 
  7   8   9  10  11  12  13 
 14  15  16  17  18  19  20 
 21  22  23  24  25  26  27 
 28  29  30  31 
```

## Calendar struct
```
type Cal struct {
	FromDate       time.Time
	ToDate         time.Time
	FirstDayOfWeek int
	MarkToday      bool
	HideHeader     bool
	ColorDefault   strings
	ColorToday     string
	ColorHighlight string
	Marker         []time.Time
}
```
#### Cal.FromDate
Specify the month to display. If no value is passed. time.Now() will be used.
NOTE: gocal will use the month - not the day. Passing 15th of August will be the same as 1st of August.
#### Cal.ToDate
Specify the last month to display. If no value is passed, only the month of Cal.FromDate is printed.
#### Cal.FirstDayOfWeek
By default, a week starts with Sunday. If you want to start with Monday (or any other day), you can set this value to 1 (for Monday).
#### Cal.MarkToday
Setting this flag to true, the todays date is highlighted in the Cal.ColorToday color.
#### Cal.HideHeader
Setting this flag to false will hide the output-header and display only the calendar without any other information.
#####Example:
```
MON TUE WED THU FRI SAT SUN
      1   2   3   4   5   6 
  7   8   9  10  11  12  13 
 14  15  16  17  18  19  20 
 21  22  23  24  25  26  27 
 28  29  30  31 
```
#### Cal.ColorDefault
Default: 29
Specify the default output color ( [ANSI Color Codes]https://en.wikipedia.org/wiki/ANSI_escape_code )
#### Cal.ColorToday
Default: 31
Specify the highlight color for today ( [ANSI Color Codes]https://en.wikipedia.org/wiki/ANSI_escape_code )
#### Cal.ColorHighlight
Default: 32
Specify the highlight color for dates passed in the Cal.Marker slice ( [ANSI Color Codes]https://en.wikipedia.org/wiki/ANSI_escape_code )
#### Cal.Marker
[]time.Time slice of dates to highlight in the calendar.
