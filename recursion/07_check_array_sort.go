package recursion

import "fmt"

func CheckIfArrayIsSortedOrNot(nums []int, index int) bool {
	fmt.Println("Index", index)
	if len(nums)-1 == index {
		return true
	}
	return nums[index] < nums[index+1] && CheckIfArrayIsSortedOrNot(nums, index+1)

}
func CheckIfArrayIsSortedOrNotHelper() {
	fmt.Println(CheckIfArrayIsSortedOrNot([]int{5, 10, 15, 20}, 0))
	fmt.Println(CheckIfArrayIsSortedOrNot([]int{6, 5, 10, 15, 20}, 0))
}
