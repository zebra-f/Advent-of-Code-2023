package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/zebra-f/Advent-of-Code-2023/day3/common"
	"github.com/zebra-f/Advent-of-Code-2023/day3/puzzle"
)

func main() {

	file, err := os.Open("day3/input.txt")
	common.ErrorCheck(err)
	defer file.Close()

	linesList := []string{}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()
	linesList = append(linesList, "." + strings.Repeat(".", len(firstLine)) + ".")
	linesList = append(linesList, "." + firstLine + ".")
	for scanner.Scan() {
		line := scanner.Text()
		linesList = append(linesList, "." + line + ".")
	}
	linesList = append(linesList, "." + strings.Repeat(".", len(firstLine)) + ".")

	puzzle.FirstPuzzle(linesList)
	puzzle.SecondPuzzle(linesList)
}