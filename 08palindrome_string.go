// ! https://leetcode.com/problems/palindrome-number/description/
package main

import "strconv"

func isIntegerPalindrome(x int) bool {
	originalString := strconv.Itoa(x)
	for i, j := 0, len(originalString)-1; i < j; i, j = i+1, j-1 {
		if originalString[i] != originalString[j] {
			return false
		}
	}
	return true
}
