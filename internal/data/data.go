package data

import (
	"os"
	"strings"
)

// Making this to get some helper functions since I'll probably have to read in a file or something

func GetData(file string) []string {
	res, err := os.ReadFile(file)
	if err != nil {
		panic("failed to read in file" + err.Error())
	}

	return strings.Split(string(res), "\n")
}
