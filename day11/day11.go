package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Galaxy struct {
	y, x int
}

func CountPaths(input []string, expansion int) int {
	galaxies := []Galaxy{}
	for y, line := range input {
		for x, point := range line {
			if point == '#' { //Found a galaxy
				galaxies = append(galaxies, Galaxy{y, x})
			}
		}
	}
	for i := len(input) - 1; i >= 0; i -= 1 {
		if isEmptyRow(input[i]) {
			for j, galaxy := range galaxies {
				if galaxy.y > i {
					galaxies[j] = Galaxy{galaxy.y + expansion - 1, galaxy.x}
				}
			}
		}
	}
	for i := len(input[0]) - 1; i >= 0; i -= 1 {
		if isEmptyCol(input, i) {
			for j, galaxy := range galaxies {
				if galaxy.x > i {
					galaxies[j] = Galaxy{galaxy.y, galaxy.x + expansion - 1}
				}
			}
		}
	}

	distance := 0
	for i := 0; i < len(galaxies)-1; i += 1 {
		for j := i + 1; j < len(galaxies); j += 1 {
			distance += int(math.Abs(float64(galaxies[i].y-galaxies[j].y)) + math.Abs(float64(galaxies[i].x-galaxies[j].x)))
		}
	}
	return distance
}

func isEmptyRow(s string) bool {
	for _, point := range s {
		if point == '#' {
			return false
		}
	}
	return true
}

func isEmptyCol(s1 []string, col int) bool {
	for _, line := range s1 {
		if line[col] == '#' {
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
	fmt.Println("Part 1:", CountPaths(input, 2))
	fmt.Println("Part 2:", CountPaths(input, 1000000))
}
