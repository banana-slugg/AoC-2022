package day1

import (
	"fmt"
	"sort"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day1/input.txt"

func Run() {
	input := help.GetSplitInput(inputFile)
	nums := make([]int, 0)

	for _, v := range input {
		nums = append(nums, help.Sum(help.MapInt(strings.Split(v, "\n"))))
	}
	sort.Ints(nums)
	fmt.Println(nums[len(nums)-1])
	fmt.Println(help.Sum(nums[len(nums)-3:]))
}
