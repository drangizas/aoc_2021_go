package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	groupsMap := map[int]int{}
	var currGroups [3]int

	// sum groups
	for i, line := range input {
		curr, _ := strconv.Atoi(line)

		var groupToPush int
		if i+3 > len(input) {
			groupToPush = 0
		} else {
			groupToPush = i + 1
		}
		currGroups = [3]int{currGroups[1], currGroups[2], groupToPush}

		for _, group := range currGroups {
			if group == 0 {
				continue
			}

			groupsMap[group] += curr
		}
	}

	// check how many increases between groups
	sortedKeys := make([]int, len(groupsMap))
	for k := range groupsMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	prev := groupsMap[sortedKeys[0]]
	for _, key := range sortedKeys {
		currVal := groupsMap[key]

		if currVal > prev {
			increases += 1
		}

		prev = currVal
	}

	fmt.Println(increases)
}
