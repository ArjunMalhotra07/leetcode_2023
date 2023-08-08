package questions

import "math"

func MaxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := int(math.Min(float64(height[left]), float64(height[right])) * float64(right-left))
		maxArea = int(math.Max(float64(maxArea), float64(area)))
		if height[left] < height[right] {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return maxArea
}