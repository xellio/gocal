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
Calendar struct:
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