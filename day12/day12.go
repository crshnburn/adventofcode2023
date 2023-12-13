package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	springLocations string
	groups          []int
}

func CountPermutations(input []string) int {
	rows := []Row{}
	for _, line := range input {
		fields := strings.Fields(line)
		groups := []int{}
		for _, valStr := range strings.Split(fields[1], ",") {
			val, _ := strconv.Atoi(valStr)
			groups = append(groups, val)
		}
		rows = append(rows, Row{fields[0], groups})
	}
	total := 0
	for _, row := range rows {
		total += count(row.springLocations, row.groups)
	}
	return total
}

func ExpandedMap(input []string) int {
	rows := []Row{}
	for _, line := range input {
		fields := strings.Fields(line)
		groups := []int{}
		for _, valStr := range strings.Split(fields[1], ",") {
			val, _ := strconv.Atoi(valStr)
			groups = append(groups, val)
		}
		totalFields := fields[0]
		totalGroups := groups
		for i := 0; i < 4; i += 1 {
			totalFields += "?" + fields[0]
			totalGroups = append(totalGroups, groups...)
		}
		rows = append(rows, Row{totalFields, totalGroups})
	}
	total := 0
	for _, row := range rows {
		total += solve(row.springLocations, row.groups)
	}
	return total
}

func count(row string, nums []int) int {
	if row == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if strings.Contains(row, "#") {
			return 0
		}
		return 1
	}

	result := 0

	if row[0] == '.' || row[0] == '?' {
		result += count(row[1:], nums)
	}

	if row[0] == '#' || row[0] == '?' {
		if nums[0] <= len(row) && !strings.Contains(row[:nums[0]], ".") && (nums[0] == len(row) || row[nums[0]] != '#') {
			if nums[0] == len(row) {
				result += count("", nums[1:])
			} else {
				result += count(row[nums[0]+1:], nums[1:])
			}
		}
	}

	return result
}

func solve(row string, groups []int) int {
	var cache [][]int
	for i := 0; i < len(row); i++ {
		cache = append(cache, make([]int, len(groups)+1))
		for j := 0; j < len(groups)+1; j++ {
			cache[i][j] = -1
		}
	}

	return dp(0, 0, row, groups, cache)
}

func dp(i, j int, record string, group []int, cache [][]int) int {
	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	res := 0
	if record[i] == '.' {
		res = dp(i+1, j, record, group, cache)
	} else {
		if record[i] == '?' {
			res += dp(i+1, j, record, group, cache)
		}
		if j < len(group) {
			count := 0
			for k := i; k < len(record); k++ {
				if count > group[j] || record[k] == '.' || count == group[j] && record[k] == '?' {
					break
				}
				count += 1
			}

			if count == group[j] {
				if i+count < len(record) && record[i+count] != '#' {
					res += dp(i+count+1, j+1, record, group, cache)
				} else {
					res += dp(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = res
	return res
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
	fmt.Println("Part 1", CountPermutations(input))
	fmt.Println("Part 2", ExpandedMap(input))
}
