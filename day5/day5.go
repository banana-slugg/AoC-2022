package day5

import (
	"fmt"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day5/input.txt"

func Part2() {
	input := help.GetSplitInput(inputFile)
	crates := strings.Split(input[0], "\n")
	crates = crates[:len(crates)-1]
	steps := strings.Split(input[1], "\n")

	parsed := make([]string, (len(crates[0])+1)/4)

	for _, v := range crates {
		for i := 0; i < (len(v)+1)/4; i++ {
			val := string(v[(i*4)+1])
			if val != " " {
				parsed[i] = val + parsed[i]
			}
		}
	}
	for _, v := range steps {
		step := strings.Split(v, " ")
		nums := help.MapIntSpread(step[1], step[3], step[5])
		amount := nums[0]
		from := (nums[1]) - 1
		to := (nums[2]) - 1
		index := help.Max(0, len(parsed[from])-amount)
		slice := parsed[from][index:]
		parsed[to] += slice
		parsed[from] = parsed[from][:index]
	}
	for _, v := range parsed {
		fmt.Print(string(v[len(v)-1]))
	}
	fmt.Println()

}

func Part1() {
	input := help.GetSplitInput(inputFile)
	crates := strings.Split(input[0], "\n")
	crates = crates[:len(crates)-1]
	steps := strings.Split(input[1], "\n")

	parsed := make([]string, (len(crates[0])+1)/4)

	for _, v := range crates {
		for i := 0; i < (len(v)+1)/4; i++ {
			val := string(v[(i*4)+1])
			if val != " " {
				parsed[i] = val + parsed[i]
			}
		}
	}
	for _, v := range steps {
		step := strings.Split(v, " ")
		nums := help.MapIntSpread(step[1], step[3], step[5])
		amount := nums[0]
		from := (nums[1]) - 1
		to := (nums[2]) - 1
		index := help.Max(0, len(parsed[from])-amount)
		slice := help.ReverseString(parsed[from][index:])
		parsed[to] += slice
		parsed[from] = parsed[from][:index]
	}
	for _, v := range parsed {
		fmt.Print(string(v[len(v)-1]))
	}
	fmt.Println()

}
