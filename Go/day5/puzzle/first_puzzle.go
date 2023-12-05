package puzzle

import (
	"fmt"
	"math"
)

func binarySearch(mapList [][3]int, target int) (middle int) {
	low := 0
	high := len(mapList) - 1
	for low <= high {
		middle = low + (high-low) / 2

		if mapList[middle][1] < target {
			low = middle + 1
		} else if mapList[middle][1] > target {
			high = middle - 1
		} else {
			return middle
		}
	}

	return low - 1
}

func FirstPuzzle(seedsIntList []int, mapNamesList [7]string, maps map[string][][3]int) (answer float64) {
	answer = math.Inf(1)
	for _, seed := range seedsIntList {
		var curr int = seed
		for _, mapName := range mapNamesList {
			index := binarySearch(maps[mapName], curr)
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
