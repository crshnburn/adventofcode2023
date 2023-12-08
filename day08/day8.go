package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"
)

type Element struct {
	left  string
	right string
}

func CountSteps(input []string) int {
	instr := input[0]
	elements := map[string]Element{}

	for i := 2; i < len(input); i += 1 {
		regexpObj := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
		matches := regexpObj.FindAllStringSubmatch(input[i], -1)
		elements[matches[0][1]] = Element{matches[0][2], matches[0][3]}
	}
	atEnd := false
	loc := "AAA"
	steps := 0
	for !atEnd {
	endLoop:
		for _, next := range instr {
			steps += 1
			element := elements[loc]
			if next == 'L' {
				loc = element.left
			} else {
				loc = element.right
			}
			if loc == "ZZZ" {
				atEnd = true
				break endLoop
			}
		}
	}
	return steps
}

func GhostSteps(input []string) int {
	instr := input[0]
	elements := map[string]Element{}

	for i := 2; i < len(input); i += 1 {
		regexpObj := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
		matches := regexpObj.FindAllStringSubmatch(input[i], -1)
		elements[matches[0][1]] = Element{matches[0][2], matches[0][3]}
	}
	locations := []string{}
	for key := range elements {
		if key[2] == 'A' {
			locations = append(locations, key)
		}
	}
	allSteps := []*big.Int{}
	for pos := range locations {
		loc := locations[pos]
		steps := int64(0)
		atEnd := false
		for !atEnd {
		endLoop:
			for _, next := range instr {
				steps += 1
				element := elements[loc]
				if next == 'L' {
					loc = element.left
				} else {
					loc = element.right
				}

				if strings.HasSuffix(loc, "Z") {
					atEnd = true
					break endLoop
				}

			}
		}
		allSteps = append(allSteps, big.NewInt(steps))
	}
	return int(lcmOfSlice(allSteps).Int64())
}

// Calculate GCD using the Euclidean algorithm
func gcd(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

// Calculate LCM of two numbers
func lcm(a, b *big.Int) *big.Int {
	if a.Sign() == 0 || b.Sign() == 0 {
		return new(big.Int)
	}
	g := gcd(a, b)
	result := new(big.Int).Div(new(big.Int).Mul(a, b), g)
	return result
}

// Calculate LCM of a slice of numbers
func lcmOfSlice(numbers []*big.Int) *big.Int {
	result := big.NewInt(1)
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
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
	fmt.Println("Part 1 Steps:", CountSteps(input))
	fmt.Println("Part 2 Steps:", GhostSteps(input))
}
