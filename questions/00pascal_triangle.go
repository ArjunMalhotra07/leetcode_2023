//! https://leetcode.com/problems/pascals-triangle/description/
package questions

func Generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j != 0 && j != i {
				row[j] = ans[i-1][j-1] + ans[i-1][j]
			} else {
				row[j] = 1
			}
		}
		ans[i] = row
	}
	return ans
}
