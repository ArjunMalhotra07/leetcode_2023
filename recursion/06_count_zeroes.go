package recursion

import "fmt"

func CountZeroes(n, numberOfZeroes int) int {
	if n == 0 {
		return numberOfZeroes
	}
	if n%10 == 0 {
		numberOfZeroes += 1
	}
	return CountZeroes(n/10, numberOfZeroes)
}
func CountZeroesHelper() {
	fmt.Println(CountZeroes(1000500, 0))
}
