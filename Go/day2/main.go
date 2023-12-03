package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/zebra-f/Advent-of-Code-2023/day2/common"
	"github.com/zebra-f/Advent-of-Code-2023/day2/puzzle"
	"github.com/zebra-f/Advent-of-Code-2023/day2/structures"
)

func main() {
	file, err := os.Open("day2/input.txt")
	common.ErrorCheck(err)
	defer file.Close()
	
	gamesList := []structures.Game{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		game := structures.Game{}
		// get Game index (the simplest way is to increment some variable by 1 in each loop)
		collonIndex := strings.Index(line, ":")
		gameIndexStr := strings.TrimSpace(line[collonIndex-2:collonIndex]) 
		gameIndexInt, err := strconv.Atoi(gameIndexStr)
		common.ErrorCheck(err)
		game.Index = gameIndexInt
		// get Game sets
		gameSets := strings.Split(line[collonIndex+1:], ";")
		gameSetsModified := []string{}
		for _, val := range gameSets {
			val = strings.ReplaceAll(val, " ", "")
			val = strings.ReplaceAll(val, "green", "g")
			val = strings.ReplaceAll(val, "blue", "b")
			val = strings.ReplaceAll(val, "red", "r")
			gameSetsModified = append(gameSetsModified, val)
		}
		game.Sets = gameSetsModified
		// append game
		gamesList = append(gamesList, game)
	}

	puzzle.FirstPuzzle(gamesList)
	puzzle.SecondPuzzle(gamesList)
}