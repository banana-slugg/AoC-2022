package day6

import (
	"fmt"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day6/input.txt"

func Part2() {
	input := help.GetInputString(inputFile)
	for i := 0; i < len(input)-3; i++ {
		temp := make(map[byte]int)

		for j := 0; j < 14; j++ {
			temp[input[i+j]]++
		}
		if len(temp) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}

func Part1() {
	input := help.GetInputString(inputFile)
	for i := 0; i < len(input)-3; i++ {
		temp := make(map[byte]int)
		for j := 0; j < 4; j++ {
			temp[input[i+j]]++
		}
		if len(temp) == 4 {
			fmt.Println(i + 4)
			break
		}
	}
}
