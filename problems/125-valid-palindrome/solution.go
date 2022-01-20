import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func isPalindrome(s string) bool {
	str := strings.Join(strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}), "")

	str = strings.ToLower(str)

	isPalindrome := true
	count := utf8.RuneCountInString(str)
	for i, j := 0, count-1; i < count; i, j = i+1, j-1 {

		if str[i] != str[j] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
