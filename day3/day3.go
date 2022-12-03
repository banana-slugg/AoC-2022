package day3

import (
	"fmt"
	"math"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day3/input.txt"

func Part2() {
	input := help.GetInput(inputFile)
	sum := 0
	for i := 0; i < len(input); i += 3 {
		freq := make(map[rune]int)
		sack := input[i]
		sacker := input[i+1]
		sackest := input[i+2]
		for _, x := range sack {
			freq[x] = 1
		}
		for _, y := range sacker {
			if freq[y] == 1 {
				freq[y]++
			}

		}
		for _, z := range sackest {
			if freq[z] == 2 {
				sum += getPriority(z)
				break
			}

		}

	}
	fmt.Println(sum)
}
func Part1() {
	input := help.GetInput(inputFile)
	sum := 0
	for _, v := range input {
		freq := make(map[rune]int)
		split := int(math.Floor(float64(len(v)) / 2))
		head := v[:split]
		tail := v[split:]
		for _, h := range head {
			if freq[h] == 0 {
				freq[h] = 1
			}
		}
		for _, t := range tail {
			if freq[t] != 0 {
				sum += getPriority(t)
				break
			}
		}
	}
	fmt.Println(sum)
}
func getPriority(c rune) int {
	if c >= 97 && c <= 122 {
		return int(c) - 96
	}
	return int(c) - 38
}
