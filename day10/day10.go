package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	North = iota
	East
	South
	West
)

const (
	None       byte = '.'
	NorthSouth byte = '|'
	EastWest   byte = '-'
	NorthEast  byte = 'L'
	NorthWest  byte = 'J'
	SouthWest  byte = '7'
	SouthEast  byte = 'F'
	Start      byte = 'S'
)

const (
	pipe = iota
	possiblyEnclosed
	open
	enclosed
	unknown
)

type position struct {
	y int
	x int
}

func CountSteps(input []string) (int, int) {
	start := position{}
	for row := range input {
		for col := range input[row] {
			if input[row][col] == Start {
				start.y = row
				start.x = col
			}
		}
	}
	steps := 0
	pos := position{}
	dir := 0
	//find connected pipe
	if start.y-1 >= 0 && (input[start.y-1][start.x] == NorthSouth || input[start.y-1][start.x] == SouthEast || input[start.y-1][start.x] == SouthWest) {
		pos.y = start.y - 1
		pos.x = start.x
		dir = North
	} else if start.x+1 < len(input[start.y]) && (input[start.y][start.x+1] == EastWest || input[start.y][start.x+1] == NorthWest || input[start.y][start.x+1] == SouthWest) {
		pos.y = start.y
		pos.x = start.x + 1
		dir = East

	} else if start.y+1 < len(input) && (input[start.y+1][start.x] == NorthSouth || input[start.y+1][start.x] == NorthEast || input[start.y+1][start.x] == NorthWest) {
		pos.y = start.y + 1
		pos.x = start.x
		dir = South

	} else if start.x-1 >= 0 && (input[start.y][start.x-1] == EastWest || input[start.y][start.x-1] == NorthEast || input[start.y][start.x-1] == SouthEast) {
		pos.y = start.y
		pos.x = start.x - 1
		dir = West
	}

	pipeLocations := make([][]byte, len(input))
	for i := range pipeLocations {
		pipeLocations[i] = make([]byte, len(input[i]))
		for j := range pipeLocations[i] {
			pipeLocations[i][j] = '.'
		}
	}
	pipeLocations[start.y][start.x] = pipe
	//while not back at the start
	for !(start.y == pos.y && start.x == pos.x) {
		pipeLocations[pos.y][pos.x] = input[pos.y][pos.x]
		switch input[pos.y][pos.x] {
		case NorthSouth:
			if dir == North {
				pos.y -= 1
			} else {
				pos.y += 1
			}
		case EastWest:
			if dir == East {
				pos.x += 1
			} else {
				pos.x -= 1
			}
		case NorthEast:
			if dir == West {
				pos.y -= 1
				dir = North
			} else {
				pos.x += 1
				dir = East
			}
		case NorthWest:
			if dir == East {
				pos.y -= 1
				dir = North
			} else {
				pos.x -= 1
				dir = West
			}
		case SouthEast:
			if dir == West {
				pos.y += 1
				dir = South
			} else {
				pos.x += 1
				dir = East
			}
		case SouthWest:
			if dir == East {
				pos.y += 1
				dir = South
			} else {
				pos.x -= 1
				dir = West
			}
		}
		steps += 1
	}

	pipeMap := make([]string, len(pipeLocations))
	for i := range pipeLocations {
		pipeMap[i] = string(pipeLocations[i])
	}

	return steps/2 + 1, FindMiddle(pipeMap)
}

func FindMiddle(input []string) int {
	count := 0
	for i := range input {
		inside := false
		wallCount := 0
		var previousWall byte
		for j := range input[i] {
			pos := input[i][j]
			if pos != '.' && pos != '-' {
				wallCount += 1
				if (pos == 'J' && previousWall == 'F') || (pos == '7' && previousWall == 'L') {
					wallCount -= 1
				}
				previousWall = pos
				if wallCount%2 == 1 {
					inside = true
				} else {
					inside = false
				}
			}
			if inside && pos == '.' {
				count += 1
			}
		}
	}
	return count
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
	steps, area := CountSteps(input)
	fmt.Println("Part 1:", steps)
	fmt.Println("Part 2:", area)
}
