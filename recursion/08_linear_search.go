package recursion

import "fmt"

func LinearSearchUsingRecursion(nums []int, index, target int) bool {
	if index == len(nums)-1 && nums[index] != target {
		return false
	}
	return nums[index] == target || LinearSearchUsingRecursion(nums, index+1, target)
}
func ReturnIndexLinearSearchUsingRecursion(nums []int, index, target int) int {
	if index == len(nums)-1 && nums[index] != target {
		return -1
	}
	if nums[index] == target {
		return index
	}
	return ReturnIndexLinearSearchUsingRecursion(nums, index+1, target)
}

func LinearSearchUsingRecursionHelper() {
	fmt.Println(LinearSearchUsingRecursion([]int{5, 10, 15, 20}, 0, 20))
	fmt.Println(LinearSearchUsingRecursion([]int{5, 10, 15, 20}, 0, 21))
	fmt.Println(ReturnIndexLinearSearchUsingRecursion([]int{5, 10, 15, 20}, 0, 20))
	fmt.Println(ReturnIndexLinearSearchUsingRecursion([]int{5, 10, 15, 20}, 0, 21))

}
