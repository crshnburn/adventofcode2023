package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CountNextElements(input []string) (int, int) {
	sequences := [][]int{}
	for _, line := range input {
		vals := []int{}
		strVals := strings.Fields(line)
		for _, strVal := range strVals {
			intVal, _ := strconv.Atoi(strVal)
			vals = append(vals, intVal)
		}
		sequences = append(sequences, vals)
	}

	totalNext := 0
	totalPrev := 0
	for _, sequence := range sequences {
		seqDiffs := [][]int{sequence}
		atEnd := false
		for !atEnd {
			diffs := findDiffs(seqDiffs[len(seqDiffs)-1])
			if allZero(diffs) {
				atEnd = true
			} else {
				seqDiffs = append(seqDiffs, diffs)
			}
		}
		nextVal := 0
		prevVal := 0
		for i := len(seqDiffs) - 1; i >= 0; i -= 1 {
			nextVal += seqDiffs[i][len(seqDiffs[i])-1]
			prevVal = seqDiffs[i][0] - prevVal
		}
		totalNext += nextVal
		totalPrev += prevVal
	}
	return totalNext, totalPrev
}

func findDiffs(sequence []int) []int {
	diffs := []int{}
	for i := 0; i < len(sequence)-1; i += 1 {
		diffs = append(diffs, sequence[i+1]-sequence[i])

	}
	return diffs
}

func allZero(sequence []int) bool {
	for _, val := range sequence {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	part1, part2 := CountNextElements(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
