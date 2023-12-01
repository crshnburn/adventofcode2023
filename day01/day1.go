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

func GetCoords(lines []string) int {
	total := 0
	for _, line := range lines {
		first, last := -1, -1
		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == -1 {
					first, _ = strconv.Atoi(string(char))
				} else {
					last, _ = strconv.Atoi(string(char))
				}
			}
		}
		if last == -1 {
			last = first
		}
		total += first*10 + last
	}
	return total
}

func GetCoords2(lines []string) int {
	pattern := "one|two|three|four|five|six|seven|eight|nine|0|1|2|3|4|5|6|7|8|9"

	regexpObj, _ := regexp.Compile(pattern)
	total := 0
	for _, line := range lines {
		matches := regexpObj.FindAllString(line, -1)
		first := matches[0]
		last := matches[len(matches)-1]
		total += ConvertMatch(first)*10 + ConvertMatch(last)
	}
	return total
}

func ConvertMatch(match string) int {
	val, err := strconv.Atoi(match)
	if err == nil {
		return val
	} else {
		switch match {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println("Coords Total:", GetCoords(lines))
	fmt.Println("Coords Second Total:", GetCoords2(lines))
}
