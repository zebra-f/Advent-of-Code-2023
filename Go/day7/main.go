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
		line = strings.Replace(line, "K", "V", -1)
		line = strings.Replace(line, "A", "Z", -1)

		handBid := strings.Split(line, " ")
		lines = append(lines, handBid)
	}
	
	puzzle.FirstPuzzle(lines)
}