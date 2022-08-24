package introduce_foreign_method

import "time"

func sendReportSol() {
	date := time.Date(2022, time.August, 10, 11, 12, 15, 15, time.UTC)
	// ...few more works
	var nextDate time.Time
	nextDate = nextDay(date)

	//...rest of the functions
	if nextDate.Before(time.Now()) {
		// ....
	}

}

func nextDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day()+1, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
}