package main

import "fmt"

func FindMargin(input [][2]int) int {
	total := 1
	for _, race := range input {
		attempts := 0
		for i := 0; i <= race[0]; i += 1 {
			hold := race[0] - i
			speed := hold
			time := race[0] - hold
			if speed*time > race[1] {
				attempts += 1
			}
		}
		total *= attempts
	}
	return total
}

func main() {
	input := [][2]int{{53, 333}, {83, 1635}, {72, 1289}, {88, 1532}}
	input2 := [][2]int{{53837288, 333163512891532}}
	fmt.Println("Part 1:", FindMargin(input))
	fmt.Println("Part 2:", FindMargin(input2))
}
