package hash

import (
	"fmt"
	"reflect"
)

// https://leetcode.cn/problems/equal-row-and-column-pairs/?envType=study-plan-v2&envId=leetcode-75

func Answer_equalPairs() {
	// First 2D array
	array1 := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}

	r1 := equalPairs(array1)
	fmt.Println(r1)

	// Second 2D array
	array2 := [][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}

	r2 := equalPairs(array2)
	fmt.Println(r2)
}

func equalPairs(grid [][]int) int {
	rowSum, columsSum := make([]int, len(grid[0])), make([]int, len(grid[0]))
	for index, value := range grid {
		for index2, value2 := range value {
			rowSum[index] += value2
			columsSum[index2] += value2
		}
	}

	var result int

	for index, row := range rowSum {

		// 看看 跟哪一列的总数相等 然后
		for index2, colum := range columsSum {
			if row != colum {
				continue
			}

			columArr := getColums(grid, index2)
			if reflect.DeepEqual(grid[index], columArr) {
				result++
			}
		}
	}
	return result
}

func getColums(grid [][]int, num int) []int {
	var result []int
	for _, value := range grid {
		result = append(result, value[num])
	}
	return result

}
