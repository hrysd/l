import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	format := strconv.Itoa(utf8.RuneCountInString(a))
	b = fmt.Sprintf("%0"+format+"s", b)

	index := len(a) - 1
	carry := 0
	result := make([]string, len(a)+1)

	for index >= 0 {
		aV, _ := strconv.Atoi(string(a[index]))
		bV, _ := strconv.Atoi(string(b[index]))

		sum := aV + bV + carry
		result = append([]string{strconv.Itoa(sum % 2)}, result...)
		carry = sum / 2

		index--
	}

	if carry > 0 {
		result = append([]string{strconv.Itoa(carry)}, result...)
	}

	return strings.Join(result, "")
}
