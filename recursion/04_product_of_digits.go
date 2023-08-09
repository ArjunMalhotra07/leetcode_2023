package recursion

import "fmt"

func ProductOfDigits(n int) int {
	if n == 0 {
		return 1
	}
	return (n % 10) * ProductOfDigits(n/10)
}

func ProductOfDigitsHelper() {
	fmt.Println(ProductOfDigits(5123))
}
