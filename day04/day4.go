package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func FindWinners(cards []string) int {
	total := 0
	for _, card := range cards {
		numbers := strings.Split(card, ":")[1]
		splitNumbers := strings.Split(numbers, "|")
		winners := strings.Split(strings.Trim(splitNumbers[0], " "), " ")
		ours := strings.Split(strings.Trim(splitNumbers[1], " "), " ")
		oursNum := []int{}
		for _, num := range ours {
			val, err := strconv.Atoi(num)
			if err == nil {
				oursNum = append(oursNum, val)
			}
		}
		found := 0
		for _, winner := range winners {
			val, err := strconv.Atoi(winner)
			if err == nil {
				for _, num := range oursNum {
					if val == num {
						found += 1
					}
				}
			}
		}
		if found > 0 {
			total += int(math.Pow(2, float64(found-1)))
		}
	}
	return total
}

func FindCardTotals(cards []string) int {
	totals := make([]int, len(cards))
	for i := range totals {
		totals[i] = 1
	}

	for cardNum, card := range cards {
		numbers := strings.Split(card, ":")[1]
		splitNumbers := strings.Split(numbers, "|")
		winners := strings.Split(strings.Trim(splitNumbers[0], " "), " ")
		ours := strings.Split(strings.Trim(splitNumbers[1], " "), " ")
		oursNum := []int{}
		for _, num := range ours {
			val, err := strconv.Atoi(num)
			if err == nil {
				oursNum = append(oursNum, val)
			}
		}
		found := 0
		for _, winner := range winners {
			val, err := strconv.Atoi(winner)
			if err == nil {
				for _, num := range oursNum {
					if val == num {
						found += 1
					}
				}
			}
		}
		for i := cardNum + 1; i <= cardNum+found; i += 1 {
			totals[i] += totals[cardNum]
		}
	}
	total := 0
	for _, val := range totals {
		total += val
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
	fmt.Println("Winning cards:", FindWinners(input))
	fmt.Println("Total cards:", FindCardTotals(input))
}
