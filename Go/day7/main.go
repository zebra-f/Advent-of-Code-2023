package main

import (
	"bufio"
	"os"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day7/puzzle"
)

func main() {
	file, err := os.Open("day7/input.txt")
	Go.ErrorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "T", "B", -1)
		line = strings.Replace(line, "J", "C", -1)
		line = strings.Replace(line, "Q", "D", -1)
		line = strings.Replace(line, "K", "E", -1)
		line = strings.Replace(line, "A", "F", -1)

		handBid := strings.Split(line, " ")
		lines = append(lines, handBid)
	}
	
	puzzle.FirstPuzzle(lines)
	puzzle.SecondPuzzle(lines)
}