package concepts

import "fmt"

func CycleSortIterative(nums []int) []int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i] - 1
		if nums[correctIndex] != nums[i] {
			//! Swap 
			temp := nums[i]
			nums[i] = nums[correctIndex]
			nums[correctIndex] = temp
		} else {
			i += 1
		}
	}
	return nums
}
func CycleSortRecursive(arr []int, index int) []int {
	if index == len(arr) {
		return arr
	}
	correctIndex := arr[index] - 1
	if arr[correctIndex] != arr[index] {
		temp := arr[index]
		arr[index] = arr[correctIndex]
		arr[correctIndex] = temp
	}
	return CycleSortRecursive(arr, index+1)
}

func CallCycleSortHelper() {
	fmt.Println(CycleSortIterative([]int{5, 2, 3, 1, 4, 6, 7, 8, 10, 9}))
	fmt.Println(CycleSortRecursive([]int{5, 2, 3, 1, 4, 6, 7, 8, 10, 9}, 0))

}

/*
Iterate from start and check if the element is at the correct Index.
If not swap it with the element which is present at that index.
If yes increment counter to check next elements.
Return array once counter reaches the last element
*/
