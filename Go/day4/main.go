package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zebra-f/Advent-of-Code-2023/day4/common"
	"github.com/zebra-f/Advent-of-Code-2023/day4/puzzle"
)


func main() {
	fmt.Println("Day 4")

	file, err := os.Open("day4/input.txt")
	common.ErrorCheck(err)
	defer file.Close()
	
	numsList := [][][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineStart := strings.Index(line, ":") + 1
		numbers := strings.Split(line[lineStart:], "|")

		var winningNums []int
		var elfsNums []int
		for _, num := range strings.Split(numbers[0], " ") {
			if num, err := strconv.Atoi(num); err == nil {
				winningNums = append(winningNums, num)
			}
		}
		for _, num := range strings.Split(numbers[1], " ") {
			if num, err := strconv.Atoi(num); err == nil {
				elfsNums = append(elfsNums, num)
			}
		}

		numsList = append(numsList, [][]int{winningNums, elfsNums})
	}
	
	puzzle.FirstPuzzle(numsList)
	puzzle.SecondPuzzle(numsList)
}