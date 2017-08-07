## gocal - a cli calendar presentation

## Motivation
Needed a calendar for my mk project - so here it is

## Usage
```
go get -u github.com/xellio/gocal
```
Print calendar:
```
calendar := gocal.Cal{}
calendar.Print()
```
## Example
### Code
```
dateLayout := "2006-01-02T15:04:05Z"
date1, _ := time.Parse(dateLayout, "2017-08-04T08:00:00Z")
date2, _ := time.Parse(dateLayout, "2017-08-14T09:30:00Z")
date3, _ := time.Parse(dateLayout, "2017-08-24T10:45:15Z")

marker := []time.Time{date1, date2, date3, date4}
fromDate := time.Now()
toDate, _ := time.Parse(dateLayout, "2017-10-24T10:45:15Z")
calendar := gocal.Cal{
	FirstDayOfWeek: 1,
	FromDate:       fromDate,
	ToDate:         toDate,
	Marker:         marker,
}
calendar.Print()
```
### Result
```
---------------------------
1 August 2017
---------------------------
MON TUE WED THU FRI SAT SUN
      1   2   3   4   5   6 
  7   8   9  10  11  12  13 
 14  15  16  17  18  19  20 
 21  22  23  24  25  26  27 
 28  29  30  31 
```

## Calendar struct:
```
type Cal struct {
	FromDate       time.Time 	// default = time.Now()
	ToDate         time.Time    // if not set - FromDate-month
	FirstDayOfWeek int 			// default sunday
	MarkToday      bool
	HideHeader     bool
	ColorDefault   strings 		// default 29
	ColorToday     string  		// default 31
	ColorHighlight string 		// default 32
	Marker         []time.Time 	// dates to mark
}
```