package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindGames(games []string) int {
	maxRed := 12
	maxBlue := 14
	maxGreen := 13
	total := 0

	for _, game := range games {
		parts := strings.Split(game, ":")
		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

		turns := strings.Split(parts[1], ";")
		valid := true
		for _, turn := range turns {
			redTotal := 0
			blueTotal := 0
			greenTotal := 0

			cubes := strings.Split(turn, ",")
			for _, cube := range cubes {
				cubeParts := strings.Split(strings.Trim(cube, " "), " ")
				count, _ := strconv.Atoi(cubeParts[0])
				switch cubeParts[1] {
				case "red":
					redTotal += count
				case "blue":
					blueTotal += count
				case "green":
					greenTotal += count
				}
			}
			if redTotal > maxRed || greenTotal > maxGreen || blueTotal > maxBlue {
				valid = false
			}
		}
		if valid {
			total += id
		}
	}
	return total
}

func FindGames2(games []string) int {
	total := 0

	for _, game := range games {
		parts := strings.Split(game, ":")

		turns := strings.Split(parts[1], ";")
		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, turn := range turns {
			cubes := strings.Split(turn, ",")
			for _, cube := range cubes {
				cubeParts := strings.Split(strings.Trim(cube, " "), " ")
				count, _ := strconv.Atoi(cubeParts[0])
				switch cubeParts[1] {
				case "red":
					maxRed = max(maxRed, count)
				case "blue":
					maxBlue = max(maxBlue, count)
				case "green":
					maxGreen = max(maxGreen, count)
				}
			}

		}
		total += maxRed * maxBlue * maxGreen
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
	var games []string
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}
	fmt.Println("Games Total:", FindGames(games))
	fmt.Println("Max Cubes:", FindGames2(games))

}
