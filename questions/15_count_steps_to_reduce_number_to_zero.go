package questions

import "fmt"

func CountStepsToConvertNumberToZero(n, numberOfSteps int) int {
	if n == 0 {
		return numberOfSteps
	}
	numberOfSteps += 1
	if n%2 == 0 {
		n /= 2
	} else {
		n -= 1
	}
	return CountStepsToConvertNumberToZero(n, numberOfSteps)
}
func CountStepsToConvertNumberToZeroHelper() {
	fmt.Println(CountStepsToConvertNumberToZero(8, 0))
}
