package utils

import (
	"fmt"
	"os"
)

func GetFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error reading file")
		return nil, err
	}
	return file, err
}

func GetKey(m map[int]int) int {
	for k := range m {
		return k
	}
	return 0
}

func GetValue(m map[int]int) int {
	for _, v := range m {
		return v
	}
	return 0
}
