package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func FindParts(input []string) int {
	pattern := "\\d+"

	regexObj := regexp.MustCompile(pattern)
	total := 0

	for lineIndex, line := range input {
		matchIndex := regexObj.FindAllStringIndex(line, -1)
		if matchIndex != nil {
			for i := 0; i < len(matchIndex); i += 1 {
			matchLoop:
				for xpos := matchIndex[i][0] - 1; xpos < matchIndex[i][1]+1; xpos += 1 {
					for ypos := lineIndex - 1; ypos <= lineIndex+1; ypos += 1 {
						if xpos >= 0 && ypos >= 0 && xpos < len(line) && ypos < len(input) {
							charat := input[ypos][xpos]
							if !(charat == '.' || unicode.IsDigit(rune(charat))) {
								value, _ := strconv.Atoi(line[matchIndex[i][0]:matchIndex[i][1]])
								total += value
								break matchLoop
							}
						}
					}
				}
			}
		}
	}
	return total
}

func FindGears(input []string) int {
	numberPattern := "\\d+"
	gearPattern := "\\*"
	total := 0

	numberRegex := regexp.MustCompile(numberPattern)
	gearRegex := regexp.MustCompile(gearPattern)

	for lineIndex, line := range input {
		gearIndex := gearRegex.FindAllStringIndex(line, -1)
		for _, gearPos := range gearIndex {
			vals := []int{}
			for j := lineIndex - 1; j <= lineIndex+1; j += 1 {
				if j >= 0 && j < len(input) {
					numberIndex := numberRegex.FindAllStringIndex(input[j], -1)
					for _, index := range numberIndex {
						if gearPos[0] >= index[0]-1 && gearPos[0] <= index[1] {
							value, _ := strconv.Atoi(input[j][index[0]:index[1]])
							vals = append(vals, value)
						}
						if len(vals) == 2 {
							total += vals[0] * vals[1]
							vals = []int{}
						}
					}
				}
			}
		}
	}
	return total
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
	fmt.Println("Part 1:", FindParts(input))
	fmt.Println("Part 2:", FindGears(input))
}
