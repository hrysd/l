func reformatDate(date string) string {
	months := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	dayRegexp := regexp.MustCompile("(th|rd|st|nd)")

	split := strings.Split(date, " ")

	year := split[2]
	month := months[split[1]]
	day, _ := strconv.Atoi(dayRegexp.ReplaceAllString(split[0], ""))

	return fmt.Sprintf("%s-%s-%02d", year, month, day)
}
