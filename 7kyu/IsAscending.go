package seventhkyu

func InAscOrder(numbers []int) bool {
	res := true

	for i, v := range numbers {
		if i+1 < len(numbers) && numbers[i+1] < v {
			res = false
		}
	}

	return res
}
