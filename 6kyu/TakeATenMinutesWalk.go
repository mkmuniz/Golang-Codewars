package kata

func IsValidWalk(walk []rune) bool {
	n := 0
	s := 0
	w := 0
	e := 0
	i := 0

	if len(walk) != 10 {
		return false
	}

	for i < len(walk) {
		switch walk[i] {
		case 'n':
			n++
			i++
		case 'w':
			w++
			i++
		case 'e':
			e++
			i++
		case 's':
			s++
			i++
		}
	}

	if n == s && e == w {
		return true
	}
	return false
}
