package main

import (
	"bufio"
	"os"

	"github.com/zebra-f/Advent-of-Code-2023/day1/puzzle"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("day1/input.txt")
	checkError(err)
	defer file.Close()

	linesList := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesList = append(linesList, line)
	}

	puzzle.FirstPuzzle(linesList)
	puzzle.SecondPuzzle(linesList)
}