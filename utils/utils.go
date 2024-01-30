package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFileScanner(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open error:", err)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
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

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringToInt(input string) int {
	output, err := strconv.Atoi(input)
	Check(err)
	return output
}
