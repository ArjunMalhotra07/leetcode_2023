package concepts

func CycleSort(nums []int) []int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i] - 1
		if nums[correctIndex] != nums[i] {
			temp := nums[i]
			nums[i] = nums[correctIndex]
			nums[correctIndex] = temp
		} else {
			i += 1
		}
	}
	return nums
}
