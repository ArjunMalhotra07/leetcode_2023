package recursion

import "fmt"

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func FactorialHelper() {
	fmt.Println(Factorial(5))
}
