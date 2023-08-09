package recursion

import "fmt"

func PrintValue(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintValue(n - 1)
	fmt.Println(n)
}

func PrintValueHelper() {
	PrintValue(5)
}
