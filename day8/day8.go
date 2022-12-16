package day8

import (
	"fmt"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day8/input.txt"

func Part2() {
	input := help.GetInput(inputFile)
	nums := make([][]int, len(input))
	for i, v := range input {
		split := strings.Split(v, "")
		nums[i] = help.MapInt(split)
	}
	scores := make([]int, 0)

	for i, row := range nums {
		for j := range row {

			scores = append(scores, getScenicScore(i, j, nums))

		}
	}

	fmt.Println(help.Max(scores...))
}

func getScenicScore(row, col int, trees [][]int) int {
	up, down, left, right := 0, 0, 0, 0
	size := len(trees[0])
	if row == 0 || col == 0 || row == size-1 || col == size-1 {
		return 0
	}

	// row forwards
	for i := row; i < size-1; i++ {
		if trees[row][col] > trees[i+1][col] {
			up++
			if i+1 == size-1 {
				break
			}
			continue
		}
		up++
		break
	}

	// row backwards
	for i := row; i > 0; i-- {
		if trees[row][col] > trees[i-1][col] {
			down++
			if i-1 == 0 {
				break
			}
			continue
		}
		down++
		break
	}

	//col forwards
	for i := col; i < size-1; i++ {
		if trees[row][col] > trees[row][i+1] {
			right++
			if i+1 == size-1 {
				break
			}
			continue
		}
		right++
		break
	}

	// col backwards
	for i := col; i > 0; i-- {
		if trees[row][col] > trees[row][i-1] {
			left++
			if i-1 == 0 {
				break
			}
			continue
		}
		left++
		break
	}

	return up * down * left * right
}

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
		}
	}

	fmt.Println(count)
}

func isVisible(row, col int, trees [][]int) bool {
	size := len(trees[0])
	if row == 0 || col == 0 || row == size-1 || col == size-1 {
		return true
	}

	// row forwards
	for i := row; i < size-1; i++ {
		if trees[row][col] > trees[i+1][col] {
			if i+1 == size-1 {
				return true
			}
			continue
		}
		break
	}

	// row backwards
	for i := row; i > 0; i-- {
		if trees[row][col] > trees[i-1][col] {
			if i-1 == 0 {
				return true
			}
			continue
		}
		break
	}

	//col forwards
	for i := col; i < size-1; i++ {
		if trees[row][col] > trees[row][i+1] {
			if i+1 == size-1 {
				return true
			}
			continue
		}
		break
	}

	// col backwards
	for i := col; i > 0; i-- {
		if trees[row][col] > trees[row][i-1] {
			if i-1 == 0 {
				return true
			}
			continue
		}
		break
	}

	return false
}
