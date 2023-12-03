package puzzle

import (
	"fmt"

	"github.com/zebra-f/Advent-of-Code-2023/day2/common"
	"github.com/zebra-f/Advent-of-Code-2023/day2/structures"
)

func FirstPuzzle(gamesList []structures.Game) int {
	answer := 0
	for _, game := range gamesList {
		red, green, blue := common.MaxColors(game.Sets)
		if red <= 12 && green <= 13 && blue <= 14 {
			answer += game.Index
		}
	}

	fmt.Println("First puzzle answer: ", answer)
	return answer
}