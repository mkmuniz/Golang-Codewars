package seventhkyu

func FindOdd(seq []int) int {
	isCompleted := false
	count := 0
	i := 0
	res := 0

	for !isCompleted {
		for _, a := range seq {
			if a == seq[i] {
				count++
			}
		}
		if count%2 == 1 {
			res = seq[i]
			isCompleted = true
		}

		i++
	}

	return res
}
