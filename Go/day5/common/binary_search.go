package common

func BinarySearch(mapList [][3]int, target int) (middle int) {
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