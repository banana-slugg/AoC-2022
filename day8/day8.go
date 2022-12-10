package day8

import (
	"fmt"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day8/input.txt"

func Part1() {
	input := help.GetInput(inputFile)
	nums := make([][]int, len(input))
	for i, v := range input {
		split := strings.Split(v, "")
		nums[i] = help.MapInt(split)
	}
	count := 0

	for i, row := range nums {
		for j := range row {

			if isVisible(i, j, nums) {
				count++
			}
			//fmt.Println(i, j, isVisible(i, j, nums), nums[i][j])
		}
	}

	fmt.Println(count)
}

func isVisible(row, col int, trees [][]int) bool {
	size := len(trees[0])
	if row == 0 || col == 0 || row == size-1 || col == size-1 {
		return true
	}

	tree := trees[row][col]
	for i := col - 1; i >= 0; i-- {
		if tree > trees[row][i] {
			if i == 0 {
				return true
			}
			tree = trees[row][i]
			continue
		}
		break
	}

	tree = trees[row][col]
	for i := col + 1; i < size; i++ {
		if tree > trees[row][i] {
			if i == size-1 {
				return true
			}
			tree = trees[row][i]
			continue
		}
		break
	}

	tree = trees[row][col]
	for i := row - 1; i >= 0; i-- {
		if tree > trees[i][col] {
			if i == 0 {
				return true
			}
			tree = trees[i][col]
			continue
		}
		break
	}

	tree = trees[row][col]
	for i := row + 1; i < size; i++ {
		if tree > trees[i][col] {
			if i == size-1 {
				return true
			}
			tree = trees[i][col]
			continue
		}
		break
	}

	return false
}
