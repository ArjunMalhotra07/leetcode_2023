// ! https://leetcode.com/problems/number-of-islands/
package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j, r := range grid[i] {
			if r == '1' {
				count++
				breadthFirstSearch(grid, i, j)
			}
		}
		fmt.Println()
	}
	return count
}
func breadthFirstSearch(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	breadthFirstSearch(grid, i-1, j)
	breadthFirstSearch(grid, i+1, j)
	breadthFirstSearch(grid, i, j+1)
	breadthFirstSearch(grid, i, j-1)
}
