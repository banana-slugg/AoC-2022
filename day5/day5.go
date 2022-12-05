package day5

import (
	"fmt"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day5/input.txt"

func Part1() {
	input := help.GetSplitInput(inputFile)
	crates := strings.Split(input[0], "\n")
	crates = crates[:len(crates)-1]
	steps := strings.Split(input[1], "\n")

	godHelpMe := make([]string, (len(crates[0])+1)/4)

	for _, v := range crates {
		for i := 0; i < (len(v)+1)/4; i++ {
			val := string(v[(i*4)+1])
			if val != " " {
				godHelpMe[i] = val + godHelpMe[i]
			}
		}
	}
	fmt.Println(godHelpMe)
	for _, v := range steps {
		step := strings.Split(v, " ")
		nums := help.MapIntSpread(step[1], step[3], step[5])
		fmt.Println(nums)
		godHelpMe[(nums[1] - 1)]
	}
}
