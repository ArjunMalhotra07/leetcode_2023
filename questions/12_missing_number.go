package questions

func MissingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i]
		if correctIndex != len(nums) && nums[correctIndex] != nums[i] {
			temp := nums[correctIndex]
			nums[correctIndex] = nums[i]
			nums[i] = temp
		} else {
			i += 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
