package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CalculateLoad(input []string) int {
	positions := [][]rune{}
	for _, line := range input {
		positions = append(positions, []rune(line))
	}
	for i := 1; i < len(positions); i += 1 {
		for pos, rock := range positions[i] {
			if rock == 'O' {
				for j := i - 1; j >= 0; j -= 1 {
					if positions[j][pos] == '.' {
						positions[j+1][pos] = '.'
						positions[j][pos] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
	// debugMap(positions)
	total := 0
	for i, line := range positions {
		lineTotal := 0
		for _, rock := range line {
			if rock == 'O' {
				lineTotal += 1
			}
		}
		total += (len(positions) - i) * lineTotal
	}

	return total
}

func CalculateLoad2(input []string) int {
	var cache = make(map[string](string))
	positions := [][]rune{}
	for _, line := range input {
		positions = append(positions, []rune(line))
	}
	start := makeKey(positions)
	for count := 0; count < 1000000000; count++ {
		key := makeKey(positions)
		if cache[key] == "" {
			for i := 1; i < len(positions); i += 1 { //north
				for pos, rock := range positions[i] {
					if rock == 'O' {
						for j := i - 1; j >= 0; j -= 1 {
							if positions[j][pos] == '.' {
								positions[j+1][pos] = '.'
								positions[j][pos] = 'O'
							} else {
								break
							}
						}
					}
				}
			}
			// debugMap(positions)
			for i := 0; i < len(positions); i += 1 { //west
				for pos, rock := range positions[i] {
					if rock == 'O' {
						for j := pos - 1; j >= 0; j -= 1 {
							if positions[i][j] == '.' {
								positions[i][j+1] = '.'
								positions[i][j] = 'O'
							} else {
								break
							}
						}
					}
				}
			}
			// debugMap(positions)
			for i := len(positions) - 1; i >= 0; i -= 1 { //south
				for pos, rock := range positions[i] {
					if rock == 'O' {
						for j := i + 1; j < len(positions); j += 1 {
							if positions[j][pos] == '.' {
								positions[j-1][pos] = '.'
								positions[j][pos] = 'O'
							} else {
								break
							}
						}
					}
				}
			}
			// debugMap(positions)
			for i := 0; i < len(positions); i += 1 { //east
				for pos := len(positions[i]) - 1; pos >= 0; pos -= 1 {
					if positions[i][pos] == 'O' {
						for j := pos + 1; j < len(positions[i]); j += 1 {
							if positions[i][j] == '.' {
								positions[i][j-1] = '.'
								positions[i][j] = 'O'
							} else {
								break
							}
						}
					}
				}
			}
			// debugMap(positions)
			cache[key] = makeKey(positions)
		} else {
			break
		}
	}
	cyclepoint := 1000000000 % len(cache)
	fmt.Println(cyclepoint)
	finalKey := start
	for i := 0; i < 1000000000; i++ {
		finalKey = cache[finalKey]
	}
	finalPos := decodeKey(finalKey, len(input[0]))
	total := 0
	for i, line := range finalPos {
		lineTotal := 0
		for _, rock := range line {
			if rock == 'O' {
				lineTotal += 1
			}
		}
		total += (len(positions) - i) * lineTotal
	}

	return total
}

func debugMap(input [][]rune) {
	for _, line := range input {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func makeKey(input [][]rune) string {
	var key string
	for _, line := range input {
		key += string(line)
	}
	return key
}

func decodeKey(input string, length int) [][]rune {
	positions := [][]rune{}

	for i := 0; i < len(input); i += length {
		end := i + length
		if end > len(input) {
			end = len(input)
		}
		positions = append(positions, []rune(input[i:end]))
	}

	return positions
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
	fmt.Println("Part 1:", CalculateLoad(input))
	fmt.Println("Part 2:", CalculateLoad2(input))
}
