package puzzle

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zebra-f/Advent-of-Code-2023/day2/structures"
)

func maxColors(gameSet []string) (red, green, blue int) {
	for _, setStr := range gameSet {
		setList := strings.Split(setStr, ",")
		for _, valColor := range setList {
			valColorLen := len(valColor)
			color := valColor[valColorLen-1:]
			val, err := strconv.Atoi(valColor[:valColorLen-1])
			if err != nil {
				panic(err)
			}
			if color == "r" {
				red = max(red, val)
			} else if color == "g" {
				green = max(green, val)
			} else {
				blue = max(blue, val)
			}
		}
	}
	return red, green, blue
}

func FirstPuzzle(gamesList []structures.Game) int {
	answer := 0
	for _, game := range gamesList {
		red, green, blue := maxColors(game.Sets)
		if red <= 12 && green <= 13 && blue <= 14 {
			answer += game.Index
		}
	}

	fmt.Println(answer)
	return answer
}