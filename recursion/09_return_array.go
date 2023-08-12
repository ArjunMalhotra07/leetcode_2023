package recursion

import "fmt"

func CheckForRepitition(nums, ans []int, index, target int) []int {
	if index == len(nums) {
		return ans
	}
	if nums[index] == target {
		ans = append(ans, index)
	}
	return CheckForRepitition(nums, ans, index+1, target)
}
func CheckForRepititionHelper() {
	fmt.Println(CheckForRepitition([]int{1, 2, 3, 10, 10, 5}, []int{}, 0, 10))
 }
