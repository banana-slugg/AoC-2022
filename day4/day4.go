package day4

import (
	"fmt"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day4/input.txt"

func Part2() {
	input := help.GetInput(inputFile)
	count := 0

	for _, v := range input {
		pair := strings.Split(v, ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")
		first_nums := help.MapInt(first)
		second_nums := help.MapInt(second)

		if first_nums[1] >= second_nums[0] && first_nums[0] <= second_nums[1] {
			count++
		}
	}
	fmt.Println(count)
}

func Part1() {
	input := help.GetInput(inputFile)
	count := 0

	for _, v := range input {
		pair := strings.Split(v, ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")
		first_nums := help.MapInt(first)
		second_nums := help.MapInt(second)

		if (first_nums[0] >= second_nums[0] && first_nums[1] <= second_nums[1]) || (second_nums[0] >= first_nums[0] && second_nums[1] <= first_nums[1]) {
			count++
		}
	}
	fmt.Println(count)
}
