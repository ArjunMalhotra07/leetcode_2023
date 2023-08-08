package concepts

import "fmt"

func SelectionSort(nums []int) []int {
	if len(nums) != 0 {
		for i := 0; i < len(nums)-1; i++ {
			smallest := nums[i]
			smallestIndex := i
			for j := i; j < len(nums); j++ {
				if nums[j] < smallest {
					smallest = nums[j]
					smallestIndex = j
				}
			}
			temp := nums[i]
			nums[i] = smallest
			nums[smallestIndex] = temp
		}
	}

	return nums
}

func CallSelectionSort() {
	// arr := []int{4, 10, 2, 3, -45, 0, 6}
	// arr1 := []int{10, -500, -5, 9, -1000, -5000, 0}
	arr2 := []int{}

	fmt.Println(SelectionSort(arr2))
}
