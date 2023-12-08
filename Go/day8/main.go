package main

import (
	"bufio"
	"fmt"
	"os"

	Go "github.com/zebra-f/Advent-of-Code-2023"
	"github.com/zebra-f/Advent-of-Code-2023/day8/puzzle"
)

func main() {
	file, err := os.Open("day8/input.txt")
	Go.ErrorCheck(err)

	scanner := bufio.NewScanner(file)

	instructions := ""
	scanner.Scan()
	instructions += scanner.Text()
	// empty line
	scanner.Scan()

	// var startPosition string
	// var endPosition string
	var instructionsMap = make(map[string][]string)
	for scanner.Scan() {
		currLine := scanner.Text()

		var currPosition, left, right string 
		currPosition = currLine[0:3]
		left, right = currLine[7:10], currLine[12:15]

		// if len(startPosition) == 0 {
		// 	startPosition = currPosition 
		// }
		// endPosition = currPosition

		instructionsMap[currPosition] = []string{left, right}	
	}
	fmt.Println(instructionsMap)

	puzzle.FirstPuzzle("AAA", "ZZZ", instructions, instructionsMap)	
}