package main

func CorrectTail(body string, tail rune) bool {
	result := rune(body[len(body)-1]) == tail

	return result
}
