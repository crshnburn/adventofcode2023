package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CalcArea(input []string) int {
	type instruction struct {
		dir    string
		amount int
		color  string
	}
	instructions := []instruction{}
	for _, line := range input {
		fields := strings.Fields(line)
		amount, _ := strconv.Atoi(fields[1])
		instructions = append(instructions, instruction{fields[0], amount, fields[2]})
	}
	type point struct {
		x, y int
	}
	currentPoint := point{0, 0}
	points := []point{currentPoint}
	perimiter := 0
	for _, instr := range instructions {
		switch instr.dir {
		case "R":
			currentPoint = point{currentPoint.x + instr.amount, currentPoint.y}
		case "L":
			currentPoint = point{currentPoint.x - instr.amount, currentPoint.y}
		case "U":
			currentPoint = point{currentPoint.x, currentPoint.y - instr.amount}
		case "D":
			currentPoint = point{currentPoint.x, currentPoint.y + instr.amount}
		}
		points = append(points, currentPoint)
		perimiter += instr.amount
	}
	n := len(points)
	area := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += points[i].x*points[j].y - points[j].x*points[i].y
	}
	if area < 0 {
		area *= -1
	}
	area /= 2
	area += (perimiter / 2) + 1
	return area
}

func CalcArea2(input []string) int {
	type instruction struct {
		dir    string
		amount int
		color  string
	}
	instructions := []instruction{}
	for _, line := range input {
		fields := strings.Fields(line)
		color := fields[2]
		lenStr := color[2:7]
		dirStr := color[7:8]
		len, _ := strconv.ParseInt(lenStr, 16, 64)
		var dir string
		switch dirStr {
		case "0":
			dir = "R"
		case "1":
			dir = "D"
		case "2":
			dir = "L"
		case "3":
			dir = "U"
		}
		instructions = append(instructions, instruction{dir, int(len), fields[2]})
	}
	type point struct {
		x, y int
	}
	currentPoint := point{0, 0}
	points := []point{currentPoint}
	perimiter := 0
	for _, instr := range instructions {
		switch instr.dir {
		case "R":
			currentPoint = point{currentPoint.x + instr.amount, currentPoint.y}
		case "L":
			currentPoint = point{currentPoint.x - instr.amount, currentPoint.y}
		case "U":
			currentPoint = point{currentPoint.x, currentPoint.y - instr.amount}
		case "D":
			currentPoint = point{currentPoint.x, currentPoint.y + instr.amount}
		}
		points = append(points, currentPoint)
		perimiter += instr.amount
	}
	n := len(points)
	area := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += points[i].x*points[j].y - points[j].x*points[i].y
	}
	if area < 0 {
		area *= -1
	}
	area /= 2
	area += (perimiter / 2) + 1
	return area
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
	fmt.Println("Part 1:", CalcArea(input))
	fmt.Println("Part 2:", CalcArea2(input))
}
