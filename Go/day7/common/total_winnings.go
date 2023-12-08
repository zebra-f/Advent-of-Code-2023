package common

import (
	"sort"
)

func GetTotalWinnings(typeMap map[int][]string, handBidMap map[string]int) (totalWinnings int) { 
	rank := 1
	for i := 0; i <= 5; i++ {
		sort.Strings(typeMap[i])
		for j := 0; j < len(typeMap[i]); j++ {
			totalWinnings += rank * handBidMap[typeMap[i][j]]
			rank += 1
		}

		if i == 3 {
			sort.Strings(typeMap[32])
			for j := 0; j < len(typeMap[32]); j++ {
				totalWinnings += rank * handBidMap[typeMap[32][j]]
				rank += 1
			}
		}
	}

	return totalWinnings
}
