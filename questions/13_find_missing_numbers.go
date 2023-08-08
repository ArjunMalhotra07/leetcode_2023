package questions

func FindDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i] - 1
		if correctIndex < len(nums) && nums[correctIndex] != nums[i] {
			temp := nums[correctIndex]
			nums[correctIndex] = nums[i]
			nums[i] = temp
		} else {
			i++
		}
	}
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
