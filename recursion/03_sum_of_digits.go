package recursion

import "fmt"

func SumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + SumOfDigits(n/10)
}

func SumOfDigitsHelper() {
	fmt.Println(SumOfDigits(1342))
}

/*
 1342 -> 1 + f(342)
				3 + f(42)
					4 + f(2)
							2+ f(0)
*/


