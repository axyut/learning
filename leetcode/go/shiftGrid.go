package main

import "fmt"

func shiftGridExecute() {
	grid := make([][]int, 0)
	grid = append(grid, []int{1, 2, 3, 4})
	grid = append(grid, []int{5, 6, 7, 8})
	grid = append(grid, []int{9, 10, 11, 12})
	fmt.Println(grid)
	fmt.Println(shiftGrid1(grid, 1))
	shiftGrid(grid, 3)
}

// very complex, performance heavy as it shifts one by one, not for all test cases.
func shiftGrid(grid [][]int, k int) [][]int {
	me, right := grid[0][0], 0

	for num := 0; num < k; num++ {
		for index, arr := range grid {
			for i := range arr {
				if i+1 >= len(arr) {

					if index+1 >= len(grid) {
						grid[0][0] = me
					} else {
						right = grid[index+1][0]
						grid[index+1][0] = me
						me = right
					}
				} else {

					right = grid[index][i+1]
					grid[index][i+1] = me
					me = right
				}

			}
		}
	}
	return grid
}

// fucking done by chatgpt
func shiftGrid1(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	total := m * n
	k = k % total

	if k == 0 {
		return grid
	}

	// Flatten the grid to a 1D array
	flat := make([]int, total)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			flat[i*n+j] = grid[i][j]
		}
	}

	// Shift the 1D array
	shifted := make([]int, total)
	for i := 0; i < total; i++ {
		shifted[(i+k)%total] = flat[i]
	}

	// Convert back to 2D grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = shifted[i*n+j]
		}
	}

	return grid
}
