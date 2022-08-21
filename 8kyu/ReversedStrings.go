package eigthkyu

import "strings"

func Solution(word string) string {
	chars := strings.Split(word, "")
	i := 1
	reversedWord := ""
	for i <= len(chars) {
		reversedWord = reversedWord + chars[len(chars)-i]
		i++
	}

	return reversedWord
}
