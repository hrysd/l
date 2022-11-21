func dayOfTheWeek(day int, month int, year int) string {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	return date.Weekday().String()
}
