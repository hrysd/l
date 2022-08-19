func numUniqueEmails(emails []string) int {
	m := map[string]int{}

	for _, email := range emails {
		splitted := strings.Split(email, "@")
		local := splitted[0]
		domain := splitted[1]
		local = strings.Split(local, "+")[0]
		local = strings.ReplaceAll(local, ".", "")

		m[local+"@"+domain] = 1
	}

	return len(m)
}
