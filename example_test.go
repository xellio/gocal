// !build +testing

package gocal_test

import (
	"time"

	"github.com/xellio/gocal"
)

//
// ExamplePrint - An example of using the Print function
//
func ExamplePrint() {
	calendar := gocal.Cal{
		FromDate: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	calendar.Print()
	// Output:
	// ---------------------------
	// January 2018
	// ---------------------------
	// SUN MON TUE WED THU FRI SAT
	//       1   2   3   4   5   6
	//   7   8   9  10  11  12  13
	//  14  15  16  17  18  19  20
	//  21  22  23  24  25  26  27
	//  28  29  30  31
}

//
// ExamplePrintWithoutHeader - Print calendar without header information
//
func ExamplePrintWithoutHeader() {
	calendar := gocal.Cal{
		FromDate:   time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		HideHeader: true,
	}
	calendar.Print()
	// Output:
	// SUN MON TUE WED THU FRI SAT
	//       1   2   3   4   5   6
	//   7   8   9  10  11  12  13
	//  14  15  16  17  18  19  20
	//  21  22  23  24  25  26  27
	//  28  29  30  31
}
