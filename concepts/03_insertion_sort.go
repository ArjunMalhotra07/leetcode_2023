package concepts

import "fmt"

func InsertionSort(nums []int) []int {
	if len(nums) != 0 {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				temp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = temp
			}
			if i+1 != 1 {
				for j := i; j > 0; j-- {
					if nums[j] < nums[j-1] {
						temp := nums[j-1]
						nums[j-1] = nums[j]
						nums[j] = temp
					}
				}
			}
		}
	}

	return nums
}

func CallInsertionSort() {
	// arr := []int{4, 10, 2, 3, -45, 0, 6}
	arr1 := []int{10, -500, -5, 9, -1000, -5000, 0}
	// arr1 := []int{7, 6, 3}

	fmt.Println(InsertionSort(arr1))
}
