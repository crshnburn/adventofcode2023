package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateHash(input []string) int {
	total := 0
	for _, line := range input {
		parts := strings.Split(line, ",")
		for _, part := range parts {
			total += calcHash(part)
		}
	}
	return total
}

func calcHash(code string) int {
	total := 0
	for _, char := range code {
		total += int(char)
		total *= 17
		total %= 256
	}
	return total
}

type Lens struct {
	label string
	power int
}

func CalcFocusingPower(input []string) int {
	total := 0
	boxes := [256][]Lens{}
	for _, line := range input {
		parts := strings.Split(line, ",")
		for _, part := range parts {
			if strings.ContainsRune(part, '-') {
				lensSplit := strings.Split(part, "-")
				hash := calcHash(lensSplit[0])
				boxes[hash] = removeLens(boxes[hash], lensSplit[0])
			} else if strings.ContainsRune(part, '=') {
				lensSplit := strings.Split(part, "=")
				hash := calcHash(lensSplit[0])
				power, _ := strconv.Atoi(lensSplit[1])
				boxes[hash] = addLens(boxes[hash], lensSplit[0], power)
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			total += (i + 1) * (j + 1) * lens.power
		}
	}

	return total
}

func removeLens(box []Lens, lens string) []Lens {
	toReturn := []Lens{}
	for _, content := range box {
		if content.label != lens {
			toReturn = append(toReturn, content)
		}
	}
	return toReturn
}

func addLens(box []Lens, lens string, power int) []Lens {
	toReturn := []Lens{}
	replaced := false
	for _, content := range box {
		if content.label == lens {
			replaced = true
			toReturn = append(toReturn, Lens{content.label, power})
		} else {
			toReturn = append(toReturn, content)
		}
	}
	if !replaced {
		toReturn = append(toReturn, Lens{lens, power})
	}
	return toReturn
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
	fmt.Println("Part 1:", CreateHash(input))
	fmt.Println("Part 2:", CalcFocusingPower(input))
}
