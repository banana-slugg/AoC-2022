package day2

import (
	"fmt"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day2/input.txt"

func Part2() {
	input := help.GetInput(inputFile)
	score := 0
	for _, v := range input {
		temp := strings.Split(v, " ")
		opp, resp := temp[0], temp[1]

		switch resp {
		case "X":
			score += makeLoss(opp)
		case "Y":
			score += 3 + makeTie(opp)
		case "Z":
			score += 6 + makeWin(opp)

		}

	}
	fmt.Println(score)
}

func Part1() {
	input := help.GetInput(inputFile)
	s := 0
	for _, v := range input {
		temp := strings.Split(v, " ")
		opp, resp := temp[0], temp[1]
		score := 0
		switch resp {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		if isWin(opp, resp) {
			score += 6
		}
		if isTie(opp, resp) {
			score += 3
		}
		s += score
	}
	fmt.Println(s)
}

func makeLoss(opp string) int {
	switch opp {
	case "A":
		return 3
	case "B":
		return 1
	case "C":
		return 2
	}
	return 0
}

func makeWin(opp string) int {
	switch opp {
	case "A":
		return 2
	case "B":
		return 3
	case "C":
		return 1
	}
	return 0
}

func makeTie(opp string) int {
	switch opp {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}
	return 0
}

func isWin(opp, resp string) bool {
	switch opp {
	case "A":
		return resp == "Y"
	case "B":
		return resp == "Z"
	case "C":
		return resp == "X"
	default:
		return false
	}
}

func isTie(opp, resp string) bool {
	switch opp {
	case "A":
		return resp == "X"
	case "B":
		return resp == "Y"
	case "C":
		return resp == "Z"
	default:
		return false
	}
}
