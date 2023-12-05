package puzzle

import (
	"fmt"
	"math"

	"github.com/zebra-f/Advent-of-Code-2023/day5/common"
)

func FirstPuzzle(seedsIntList []int, mapNamesList [7]string, maps map[string][][3]int) (answer float64) {
	answer = math.Inf(1)
	for _, seed := range seedsIntList {
		var curr int = seed
		for _, mapName := range mapNamesList {
			index := common.BinarySearch(maps[mapName], curr)
			if index == -1 {
				continue
			}
			var destination, source, range_ = maps[mapName][index][0], maps[mapName][index][1], maps[mapName][index][2]
			if curr > source + range_ {
				continue
			} else {
				diff := curr - source
				curr = destination + diff
			}
		}
		answer = min(answer, float64(curr))
	}

	fmt.Println("First puzzle answer: ", int(answer))
	return answer
}
