// ! https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
package main

func twoSum(numbers []int, target int) []int {
	s := 0
	e := len(numbers) - 1
	for s <= e {
		sum := numbers[s] + numbers[e]
		if sum > target {
			e -= 1
		} else if sum < target {
			s += 1
		} else {
			return []int{s + 1, e + 1}
		}
	}
	return []int{}
}
