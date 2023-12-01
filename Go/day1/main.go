package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zebra-f/Advent-of-Code-2023/day1/puzzle"
)
func check_error(err error) {
	if err != nil {
		panic(err)

	}
}

func main() {
	fmt.Println("Test Day1")
	file, err := os.Open("day1/input.txt")
	check_error(err)
	defer file.Close()

	linesList := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesList = append(linesList, line)
	}

	puzzle.FirstPuzzle(linesList)
}