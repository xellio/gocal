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
	FromDate       time.Time    // default = time.Now()
	ToDate         time.Time    // if not set - FromDate-month
	FirstDayOfWeek int          // default 0 (sunday)
	MarkToday      bool         // default false
	HideHeader     bool         // default false
	ColorDefault   strings      // default 29
	ColorToday     string       // default 31
	ColorHighlight string       // default 32
	Marker         []time.Time  // dates to mark
}
```