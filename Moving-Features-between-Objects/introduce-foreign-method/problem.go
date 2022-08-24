package introduce_foreign_method

import "time"

func sendReport() {
	previousDate := time.Date(2022, time.August, 10, 11, 12, 15, 15, time.UTC)
	// ...few more works
	var nextDate time.Time
	nextDate = time.Date(previousDate.Year(), previousDate.Month(), previousDate.Day()+1, previousDate.Hour(), previousDate.Minute(), previousDate.Second(), previousDate.Nanosecond(), previousDate.Location())

	//...rest of the functions
	if nextDate.Before(time.Now()) {
		// ....
	}

}
