package eigthkyu

import "fmt"

func countSheep(num int) string {
	res := ""
	i := 1
	for i <= num {
		res = res + fmt.Sprint(i) + " sheep..."
		i++
	}
	return res
}
