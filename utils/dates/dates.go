package dates

import "time"

func Today() (time.Time, time.Time) {
	year, month, day := time.Now().Date()
	dateStart := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

	dateEnd := dateStart.Add(time.Hour*time.Duration(23) +
		time.Minute*time.Duration(59) +
		time.Second*time.Duration(59))

	return dateStart, dateEnd
}

func LastWeek() (time.Time, time.Time) {
	now := time.Now()
	year, month, day := now.Date()

	dateStart := time.Date(year, month, day, 0, 0, 0, 0, now.Location()).AddDate(0, 0, -7)

	dateEnd := dateStart.AddDate(0, 0, 7).Add(time.Hour*time.Duration(23) +
		time.Minute*time.Duration(59) +
		time.Second*time.Duration(59))

	return dateStart, dateEnd
}

func Of(date time.Time) (time.Time, time.Time) {
	dateStart := date
	dateEnd := date.Add(time.Hour*time.Duration(23) +
		time.Minute*time.Duration(59) +
		time.Second*time.Duration(59))

	return dateStart, dateEnd
}

func Between(startDate time.Time, endDate time.Time) (time.Time, time.Time) {

	dateStart := time.Date(
		startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	dateEnd := time.Date(
		endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, endDate.Location())

	return dateStart, dateEnd
}
