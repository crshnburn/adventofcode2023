package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CountReflections(input []string) int {
	maps := [][]string{}
	currentMap := []string{}
	for _, line := range input {
		if len(line) == 0 {
			maps = append(maps, currentMap)
			currentMap = []string{}
		} else {
			currentMap = append(currentMap, line)
		}
	}
	maps = append(maps, currentMap)
	total := 0
	for _, currentMap := range maps {
		for i := 0; i < len(currentMap[0])-1; i += 1 {
			if hasVerticalSymmetry(currentMap, i) {
				// fmt.Println("Map: ", mapNo, "Score:", i+1)
				// debugMap(currentMap, i, true)
				total += i + 1
				break
			}
		}
		for i := 0; i < len(currentMap)-1; i += 1 {
			if hasHorizontalSymmetry(currentMap, i) {
				// fmt.Println("Map: ", mapNo, "Score:", i+1)
				// debugMap(currentMap, i, false)
				total += 100 * (i + 1)
				break
			}
		}
	}
	return total
}

func CountReflectionsWithSmudge(input []string) int {
	maps := [][]string{}
	currentMap := []string{}
	for _, line := range input {
		if len(line) == 0 {
			maps = append(maps, currentMap)
			currentMap = []string{}
		} else {
			currentMap = append(currentMap, line)
		}
	}
	maps = append(maps, currentMap)
	total := 0

	for mapNo, currentMap := range maps {
		found := false
		origVert := -1
		origHoriz := -1
		for i := 0; i < len(currentMap[0])-1; i += 1 {
			if hasVerticalSymmetry(currentMap, i) {
				// fmt.Println("Map: ", mapNo, "Score:", i+1)
				// debugMap(currentMap, i, true)
				origVert = i
				break
			}
		}
		for i := 0; i < len(currentMap)-1; i += 1 {
			if hasHorizontalSymmetry(currentMap, i) {
				// fmt.Println("Map: ", mapNo, "Score:", i+1)
				// debugMap(currentMap, i, false)
				origHoriz = i
				break
			}
		}

		for y := 0; y < len(currentMap); y += 1 {
			for x := 0; x < len(currentMap[y]); x += 1 {
				smudgedMap := []string{}
				for i, line := range currentMap {
					if i == y {
						smudgedMap = append(smudgedMap, smudgeLine(line, x))
					} else {
						smudgedMap = append(smudgedMap, line)
					}
				}

				for i := 0; i < len(smudgedMap[0])-1; i += 1 {
					if hasVerticalSymmetry(smudgedMap, i) && origVert != i {
						fmt.Println("Map: ", mapNo, "Score:", i+1)
						fmt.Println("Smudge at", y, x, "Index", i)
						debugMap(smudgedMap, i, true)
						total += i + 1
						found = true
						break
					}
				}
				for i := 0; i < len(smudgedMap)-1; i += 1 {
					if hasHorizontalSymmetry(smudgedMap, i) && origHoriz != i {
						fmt.Println("Map: ", mapNo, "Score:", (i+1)*100)
						fmt.Println("Smudge at", y, x)
						debugMap(smudgedMap, i, false)
						total += 100 * (i + 1)
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
	}

	return total
}

func smudgeLine(line string, point int) string {
	runes := []rune(line)
	for i, char := range line {
		if i == point {
			if char == '#' {
				runes[i] = '.'
			} else if char == '.' {
				runes[i] = '#'
			}
		}
	}
	return string(runes)
}

func checkVertialSymmetry(currentMap []string, pos1, pos2 int) bool {
	for i := 0; i < len(currentMap); i += 1 {
		if currentMap[i][pos1] != currentMap[i][pos2] {
			return false
		}
	}
	return true
}

func hasVerticalSymmetry(currentMap []string, start int) bool {
	col1 := start
	col2 := start + 1
	for col1 >= 0 && col2 < len(currentMap[0]) {
		if checkVertialSymmetry(currentMap, col1, col2) {
			col1 -= 1
			col2 += 1
		} else {
			return false
		}
	}
	return true
}

func hasHorizontalSymmetry(currentMap []string, start int) bool {
	row1 := start
	row2 := start + 1
	for row1 >= 0 && row2 < len(currentMap) {
		if currentMap[row1] == currentMap[row2] {
			row1 -= 1
			row2 += 1
		} else {
			return false
		}
	}
	return true
}

func debugMap(input []string, pos int, isVertical bool) {
	if isVertical {
		for i := 0; i < pos; i += 1 {
			fmt.Print(" ")
		}
		fmt.Println("V")
		for _, line := range input {
			fmt.Println(line)
		}
	} else {
		for i, line := range input {
			if i == pos {
				fmt.Println("> ", line)
			} else {
				fmt.Println("  ", line)
			}
		}
	}
	fmt.Println()
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
	fmt.Println("Part 1:", CountReflections(input))
	fmt.Println("Part 2:", CountReflectionsWithSmudge(input))
}
