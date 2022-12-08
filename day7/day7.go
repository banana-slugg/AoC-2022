package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/crims1n/AoC-2022/help"
)

const inputFile = "./day7/input.txt"

func Part2() {
	input := help.GetInput(inputFile)
	dirs := make(map[string]int)
	path := make([]string, 0)
	for _, v := range input {
		cmd := strings.Split(v, " ")
		if cmd[1] == "ls" {
			continue
		}
		if cmd[1] == "cd" {
			switch cmd[2] {
			case "..":
				path = path[:len(path)-1]
			default:
				path = append(path, cmd[2])
			}
			continue
		}

		num, err := strconv.Atoi(cmd[0])
		if err != nil {
			continue
		}
		for i := range path {
			dirs[getWorkingDir(path[:i+1])] += num
		}

	}
	valid := make([]int, 0)
	freeSpace := 70000000 - dirs["/"]
	spaceNeeded := 30000000 - freeSpace
	for _, v := range dirs {
		if (spaceNeeded - v) <= 0 {
			valid = append(valid, v)
		}
	}
	sort.Ints(valid)
	fmt.Println(valid[0])
}

func Part1() {
	input := help.GetInput(inputFile)
	dirs := make(map[string]int)
	path := make([]string, 0)
	for _, v := range input {
		cmd := strings.Split(v, " ")
		if cmd[1] == "ls" {
			continue
		}
		if cmd[1] == "cd" {
			switch cmd[2] {
			case "..":
				path = path[:len(path)-1]
			default:
				path = append(path, cmd[2])
			}
			continue
		}

		num, err := strconv.Atoi(cmd[0])
		if err != nil {
			continue
		}
		for i := range path {
			dirs[getWorkingDir(path[:i+1])] += num
		}

	}
	sum := 0
	for _, v := range dirs {
		if v <= 100000 {
			sum += v
		}
	}
	fmt.Println(sum)
}

func getWorkingDir(dirs []string) string {
	temp := ""
	for _, v := range dirs {
		temp += v
	}
	return temp
}
