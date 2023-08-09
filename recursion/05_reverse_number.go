package recursion

import "fmt"

func ReverseNumber(n, reversedNumber int) int {
	if n%10 == 0 {
		return reversedNumber
	}
	rem := n % 10
	return ReverseNumber(n/10, reversedNumber*10+rem)
}
func ReverseNumberHelper() {
	fmt.Println(ReverseNumber(2134, 0))
}
