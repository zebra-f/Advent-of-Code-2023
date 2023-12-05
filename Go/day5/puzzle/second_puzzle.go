package puzzle

import (
	"fmt"
	"math"

	"github.com/zebra-f/Advent-of-Code-2023/day5/common"
)

func SecondPuzzle(seedsIntList []int, mapNamesList [7]string, maps map[string][][3]int) (answer float64) {
	answer = math.Inf(1)
	for i, seed := range seedsIntList {
		if i % 2 != 0 || i == len(seedsIntList) - 1{
			continue
		}
		for j := seed; j < seed + seedsIntList[i+1]; j++ {
			var curr int = j
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
	}

	fmt.Println("Second puzzle answer: ", int(answer))
	return answer
}