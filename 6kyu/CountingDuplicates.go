package kata

import "strings"

func duplicate_count(s string) int {
	a := 0
	counter := 0
	letterCounter := 0
	str := strings.ToLower(s)
	usedLetters := ""

	for a < len(str) {
		if strings.Count(usedLetters, string(str[a])) == 0 {
			counter = strings.Count(str, string(str[a]))
			usedLetters = usedLetters + string(str[a])
		}
		usedLetters = usedLetters + string(str[a])
		if counter > 1 {
			counter = 0
			letterCounter++
			a++
		}
		counter = 0
		a++
	}

	return letterCounter
}
