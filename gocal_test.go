package gocal

import (
	"testing"
	"time"
)

type testpairInitDate struct {
	cal      *Cal
	expected *Cal
}

var testCasesInitDate = []testpairInitDate{
	{
		testInitDateBuildCal("2017-09-11T11:11:11Z", ""),
		testInitDateBuildCal("2017-09-01T00:00:00Z", "2017-09-30T00:00:00Z"),
	},
}

func testInitDateBuildCal(fromDate string, toDate string) *Cal {

	fd, _ := time.Parse("2006-01-02T15:04:05Z", fromDate)
	td, _ := time.Parse("2006-01-02T15:04:05Z", toDate)

	cal := new(Cal)
	cal.FromDate = fd
	cal.ToDate = td
	return cal
}

func TestInitDate(t *testing.T) {
	for i, pair := range testCasesInitDate {
		pair.cal.init()
		if pair.cal.FromDate != pair.expected.FromDate {
			t.Error(
				"\nFromDate:",
				"\nGot", pair.cal.FromDate,
				"\nExpected", pair.expected.FromDate,
				"\nTestCase", i,
			)
		}

		if pair.cal.ToDate != pair.expected.ToDate {
			t.Error(
				"\nToDate:",
				"\nGot", pair.cal.ToDate,
				"\nExpected", pair.expected.ToDate,
				"\nTestCase", i,
			)
		}

	}
}
