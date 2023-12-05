package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day5/puzzle"
)

func main() {
	
	file, err := os.Open("day5/input.txt")
	Go.ErrorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	
	// get seeds list
	seedsLine := scanner.Text()
	seedsStart := strings.Index(seedsLine, ":")
	seedsStrList := strings.Split(seedsLine[seedsStart+2:], " ")

	var seedsIntList []int
	for _, val := range seedsStrList {
		if intVal, err := strconv.Atoi(val); err == nil {
			seedsIntList = append(seedsIntList, intVal)
		}
	}
	
	// empty line
	scanner.Scan()
	
	// map destination/source maps
	maps := make(map[string][][3]int)	
	mapNamesList := [7]string{
		"seed-to-soil map:", 
		"soil-to-fertilizer map:", 
		"fertilizer-to-water map:", 
		"water-to-light map:", 
		"light-to-temperature map:", 
		"temperature-to-humidity map:", 
		"humidity-to-location map:",
	}
	
	var mapCounter int = -1
	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			continue
		}
		for _, mapName := range mapNamesList {
			if currLine == mapName {
				mapCounter += 1
				// move to the next line
				scanner.Scan()
				currLine = scanner.Text()
				break

			}
		}

		// convert string ranges to int
		seedsRangeStrList := strings.Split(currLine, " ")
		seedsRangeIntList := [3]int{}
		for i, val := range seedsRangeStrList {
			if intVal, err := strconv.Atoi(val); err == nil {
				seedsRangeIntList[i] = intVal
			}
		}
		maps[mapNamesList[mapCounter]] = append(maps[mapNamesList[mapCounter]], seedsRangeIntList) 

	}

	// sort by source
	for key := range maps {
		sort.Slice(maps[key], func (i, j int) bool {
			return maps[key][i][1] < maps[key][j][1]
		})
	}

	puzzle.FirstPuzzle(seedsIntList, mapNamesList, maps)
}