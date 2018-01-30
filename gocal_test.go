// !build +testing

package gocal

import (
	"testing"
	"time"
)

type testpair struct {
	cal      *Cal
	expected *Cal
}

var testCasesInitDate = []testpair{
	{
		buildCalForTestInitDate("2017-09-11T11:11:11Z", ""),
		buildCalForTestInitDate("2017-09-01T00:00:00Z", "2017-09-30T00:00:00Z"),
	},
}

func buildCalForTestInitDate(fromDate string, toDate string) *Cal {

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
				t.Name(),
				"\nFromDate:",
				"\nGot", pair.cal.FromDate,
				"\nExpected", pair.expected.FromDate,
				"\nTestCase", i,
			)
		}

		if pair.cal.ToDate != pair.expected.ToDate {
			t.Error(
				t.Name(),
				"\nToDate:",
				"\nGot", pair.cal.ToDate,
				"\nExpected", pair.expected.ToDate,
				"\nTestCase", i,
			)
		}

	}
}

type testpairShouldBeMarked struct {
	cal        *Cal
	dateToMark string
	expected   bool
}

var testCasesShouldBeMarked = []testpairShouldBeMarked{
	{
		buildCalForTestShouldBeMarked([]string{}),
		"",
		false,
	},
	{
		buildCalForTestShouldBeMarked([]string{}),
		"2017-08-14T09:30:00Z",
		false,
	},
	{
		buildCalForTestShouldBeMarked([]string{"2017-08-04T08:00:00Z", "2017-08-14T09:30:00Z", "2017-08-24T10:45:15Z", "2017-09-24T10:45:15Z"}),
		"2017-08-14T09:30:00Z",
		true,
	},
	{
		buildCalForTestShouldBeMarked([]string{"2017-08-04T08:00:00Z"}),
		"2017-08-04T18:00:00Z",
		true,
	},
	{
		buildCalForTestShouldBeMarked([]string{"2017-08-04T08:00:00Z", "2017-08-14T09:30:00Z", "2017-08-24T10:45:15Z", "2017-09-24T10:45:15Z"}),
		"2018-08-14T09:30:00Z",
		false,
	},
}

func buildCalForTestShouldBeMarked(markerDates []string) *Cal {

	var marker []time.Time
	for _, d := range markerDates {
		date, _ := time.Parse("2006-01-02T15:04:05Z", d)
		marker = append(marker, date)
	}

	cal := new(Cal)
	cal.Marker = marker

	return cal
}

func TestShouldBeMarked(t *testing.T) {
	for i, testcase := range testCasesShouldBeMarked {
		dateToMark, _ := time.Parse("2006-01-02T15:04:05Z", testcase.dateToMark)
		if mark := testcase.cal.shouldBeMarked(dateToMark); mark != testcase.expected {
			t.Error(
				t.Name(),
				"\nGot", mark,
				"\nExpected", testcase.expected,
				"\nTestCase", i,
			)
		}
	}
}

func TestCalculateWeeks(t *testing.T) {
	t.SkipNow()
}
