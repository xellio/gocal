package gocal

var (
	MONDAY    = "MON"
	TUESDAY   = "TUE"
	WEDNESDAY = "WED"
	THURSDAY  = "THU"
	FRIDAY    = "FRI"
	SATURDAY  = "SAT"
	SUNDAY    = "SUN"

	DAYS = map[string]int{
		SUNDAY:    0,
		MONDAY:    1,
		TUESDAY:   2,
		WEDNESDAY: 3,
		THURSDAY:  4,
		FRIDAY:    5,
		SATURDAY:  6,
	}
	FIRST_DAY_OF_WEEK = 1
)
