package help

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Making this to get some helper functions since I'll probably have to read in a file or something

func GetInput(file string) []string {
	res, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(res), "\n")
}

func GetSplitInput(file string) []string {
	res, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(res), "\n\n")
}

func GetInputString(file string) string {
	res, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(res))
}

func MapInt(input []string) []int {
	temp := make([]int, len(input))
	for i, v := range input {
		converted, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		temp[i] = converted
	}
	return temp
}

func MapIntSpread(input ...string) []int {
	temp := make([]int, len(input))
	for i, v := range input {
		converted, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		temp[i] = converted
	}
	return temp
}
