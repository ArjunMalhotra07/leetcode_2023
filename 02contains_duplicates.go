//! https://leetcode.com/problems/contains-duplicate/ 
package main

func containsDuplicate(nums []int) bool {
	ans := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exist := ans[nums[i]]; exist {
			return true
		} else {
			ans[nums[i]] = 1
		}
	}
	return false
}
