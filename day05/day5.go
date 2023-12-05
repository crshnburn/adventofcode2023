package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type AlmanacMap struct {
	destinationStart int
	sourceStart      int
	length           int
}

func findLoc(initialLoc int, locMap []AlmanacMap) int {
	locVal := initialLoc
	for _, loc := range locMap {
		if initialLoc >= loc.sourceStart && initialLoc < loc.sourceStart+loc.length {
			locVal = loc.destinationStart + initialLoc - loc.sourceStart
		}
	}
	return locVal
}

func FindLowestLocation(lines []string) int {
	seeds := ""
	seedToSoil := []AlmanacMap{}
	soilToFertilizer := []AlmanacMap{}
	fertilizerToWater := []AlmanacMap{}
	waterToLight := []AlmanacMap{}
	lightToTemperature := []AlmanacMap{}
	temperatureToHumidity := []AlmanacMap{}
	humidityToLocation := []AlmanacMap{}

	var currentMap *[]AlmanacMap
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seeds = strings.Trim(strings.Split(line, ":")[1], " ")
		} else if line == "seed-to-soil map:" {
			currentMap = &seedToSoil
		} else if line == "soil-to-fertilizer map:" {
			currentMap = &soilToFertilizer
		} else if line == "fertilizer-to-water map:" {
			currentMap = &fertilizerToWater
		} else if line == "water-to-light map:" {
			currentMap = &waterToLight
		} else if line == "light-to-temperature map:" {
			currentMap = &lightToTemperature
		} else if line == "temperature-to-humidity map:" {
			currentMap = &temperatureToHumidity
		} else if line == "humidity-to-location map:" {
			currentMap = &humidityToLocation
		} else if len(line) > 0 {
			//Must be a list of values
			var newMap AlmanacMap
			fmt.Sscanf(line, "%d %d %d", &newMap.destinationStart, &newMap.sourceStart, &newMap.length)
			*currentMap = append(*currentMap, newMap)
		}
	}

	lowestLocation := math.MaxInt
	for _, seed := range strings.Split(seeds, " ") {
		seedVal, _ := strconv.Atoi(seed)
		soilVal := findLoc(seedVal, seedToSoil)
		fertlizerVal := findLoc(soilVal, soilToFertilizer)
		waterVal := findLoc(fertlizerVal, fertilizerToWater)
		lightVal := findLoc(waterVal, waterToLight)
		tempVal := findLoc(lightVal, lightToTemperature)
		humidityVal := findLoc(tempVal, temperatureToHumidity)
		locationVal := findLoc(humidityVal, humidityToLocation)
		if locationVal < lowestLocation {
			lowestLocation = locationVal
		}
	}
	return lowestLocation
}

func FindLowestLocationRange(lines []string) int {
	seeds := ""
	seedToSoil := []AlmanacMap{}
	soilToFertilizer := []AlmanacMap{}
	fertilizerToWater := []AlmanacMap{}
	waterToLight := []AlmanacMap{}
	lightToTemperature := []AlmanacMap{}
	temperatureToHumidity := []AlmanacMap{}
	humidityToLocation := []AlmanacMap{}

	var currentMap *[]AlmanacMap
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seeds = strings.Trim(strings.Split(line, ":")[1], " ")
		} else if line == "seed-to-soil map:" {
			currentMap = &seedToSoil
		} else if line == "soil-to-fertilizer map:" {
			currentMap = &soilToFertilizer
		} else if line == "fertilizer-to-water map:" {
			currentMap = &fertilizerToWater
		} else if line == "water-to-light map:" {
			currentMap = &waterToLight
		} else if line == "light-to-temperature map:" {
			currentMap = &lightToTemperature
		} else if line == "temperature-to-humidity map:" {
			currentMap = &temperatureToHumidity
		} else if line == "humidity-to-location map:" {
			currentMap = &humidityToLocation
		} else if len(line) > 0 {
			//Must be a list of values
			var newMap AlmanacMap
			fmt.Sscanf(line, "%d %d %d", &newMap.destinationStart, &newMap.sourceStart, &newMap.length)
			*currentMap = append(*currentMap, newMap)
		}
	}

	lowestLocation := math.MaxInt
	var seedRanges [][2]int
	seedFields := strings.Fields(seeds)
	for i := 0; i < len(seedFields); i += 2 {
		start, _ := strconv.Atoi(seedFields[i])
		length, _ := strconv.Atoi(seedFields[i+1])
		seedRanges = append(seedRanges, [2]int{start, length})
	}
	for _, seedRange := range seedRanges {
		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]; seed += 1 {
			soilVal := findLoc(seed, seedToSoil)
			fertlizerVal := findLoc(soilVal, soilToFertilizer)
			waterVal := findLoc(fertlizerVal, fertilizerToWater)
			lightVal := findLoc(waterVal, waterToLight)
			tempVal := findLoc(lightVal, lightToTemperature)
			humidityVal := findLoc(tempVal, temperatureToHumidity)
			locationVal := findLoc(humidityVal, humidityToLocation)
			if locationVal < lowestLocation {
				lowestLocation = locationVal
			}
		}
	}
	return lowestLocation
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
	fmt.Println("Lowest location:", FindLowestLocation(input))
	fmt.Println("Lowest location range:", FindLowestLocationRange(input))
}
