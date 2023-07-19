// ! https://leetcode.com/problems/number-of-islands/
package questions

import "fmt"

func NumIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j, r := range grid[i] {
			if r == '1' {
				count++
				BreadthFirstSearch(grid, i, j)
			}
		}
		fmt.Println()
	}
	return count
}
func BreadthFirstSearch(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	BreadthFirstSearch(grid, i-1, j)
	BreadthFirstSearch(grid, i+1, j)
	BreadthFirstSearch(grid, i, j+1)
	BreadthFirstSearch(grid, i, j-1)
}
