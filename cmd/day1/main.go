package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/crims1n/AoC-2022/internal/data"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	p, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}
	path := filepath.Dir(p)
	input := data.GetData(path + "/input.txt")
	temp := 0
	ans := 0
	big := make([]int, 0)
	for _, v := range input {
		if len(v) != 0 {
			val, _ := strconv.Atoi(v)
			temp += val
		} else {
			if temp > ans {
				ans = temp
				big = append(big, ans)
			}
			temp = 0
		}
	}
	sort.Ints(big)
	biggest := big[len(big)-3:]
	blah := 0

	for _, v := range biggest {
		blah += v
	}
	fmt.Println(blah)
}

func partOne() {
	p, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}
	path := filepath.Dir(p)
	input := data.GetData(path + "/input.txt")
	temp := 0
	ans := 0
	for _, v := range input {
		if len(v) != 0 {
			val, _ := strconv.Atoi(v)
			temp += val
		} else {
			if temp > ans {
				ans = temp
			}
			temp = 0
		}
	}

	fmt.Println(ans)
}
