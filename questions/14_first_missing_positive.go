package questions

func FirstMissingPositive(nums []int) int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i] - 1
		if nums[i] > 0 && nums[i] <= len(nums) && nums[correctIndex] != nums[i] {
			temp := nums[correctIndex]
			nums[correctIndex] = nums[i]
			nums[i] = temp
		} else {
			i += 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
