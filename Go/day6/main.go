package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day6/puzzle"
)

func main() {
	file, err := os.Open("day6/input.txt")
	Go.ErrorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeLine := scanner.Text()
	var times []int
	for _, value := range strings.Split(timeLine, " ") {
		if intValue, err := strconv.Atoi(value); err == nil {
			times = append(times, intValue)
		}
	}

	scanner.Scan()
	distanceLine := scanner.Text()
	var distances []int
	for _, value := range strings.Split(distanceLine, " ") {
		if intValue, err := strconv.Atoi(value); err == nil {
			distances = append(distances, intValue)
		}
	}

	puzzle.FirstPuzzle(times, distances)

}