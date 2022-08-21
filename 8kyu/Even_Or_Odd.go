package eigthkyu

import "fmt"

func main() {
	fmt.Println(EvenOrOdd(2))
	fmt.Println(EvenOrOdd(3))
}

func EvenOrOdd(number int) string {
	result := number % 2

	if result == 0 {
		return "Even"
	}

	return "Odd"
}
