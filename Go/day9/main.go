package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day9/puzzle"
)

func main() {
	file, err := os.Open("day9/input.txt")
	Go.ErrorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]int{}
	for scanner.Scan() {
		currLine := scanner.Text()

		line := strings.Split(currLine, " ")
		intLine := []int{}
		for _, strNum := range line {
			num, err := strconv.Atoi(strNum)
			Go.ErrorCheck(err)
			intLine = append(intLine, num)
		}

		lines = append(lines, intLine)
	}

	puzzle.FirstPuzzle(lines)
	puzzle.SecondPuzzle(lines)
}