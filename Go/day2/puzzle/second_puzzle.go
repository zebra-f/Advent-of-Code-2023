package puzzle

import (
	"fmt"

	"github.com/zebra-f/Advent-of-Code-2023/day2/common"
	"github.com/zebra-f/Advent-of-Code-2023/day2/structures"
)

func SecondPuzzle(gamesList []structures.Game) int {
	answer := 0
	for _, game := range gamesList {
		red, green, blue := common.MaxColors(game.Sets)
		answer += red * blue * green
	}

	fmt.Println(answer)
	return answer
}