package concepts

import "fmt"

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		didSwap := false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
				didSwap = true
			}
		}
		if !didSwap {
			return nums
		}
	}
	return nums
}
func BubbleSortRecursive(nums []int, index, pass int) []int {
	if pass == 0 {
		return nums
	}
	if index < pass {
		if nums[index] > nums[index+1] {
			temp := nums[index]
			nums[index] = nums[index+1]
			nums[index+1] = temp
		}
		return BubbleSortRecursive(nums, index+1, pass)
	} else {
		return BubbleSortRecursive(nums, 0, pass-1)
	}
}

func CallBubbleSort() {
	arr1 := []int{4, 10, 2, 3, -45, 0, 6}
	// arr1 := []int{1, 2, 3, 4, 5}

	fmt.Println(BubbleSort(arr1))
	fmt.Println(BubbleSortRecursive(arr1, 0, 6))

}
