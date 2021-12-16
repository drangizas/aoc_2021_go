package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFileIntoStringArray(fileName string) []string {
	var lines []string

	input, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	input.Close()

	return lines
}

func main() {
	input := readFileIntoStringArray("day1/day1.input.txt")
	increases := 0

	prev, _ := strconv.Atoi(input[0])
	for i := 1; i < len(input); i++ {
		curr, _ := strconv.Atoi(input[i])

		if curr > prev {
			increases += 1
		}

		prev = curr
	}

	fmt.Println(increases)
}
