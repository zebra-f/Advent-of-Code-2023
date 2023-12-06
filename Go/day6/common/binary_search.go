package common

func LowerBoundBinarySearch(low, high, target int) int {
	time := high	
	for low < high {
		middle := low + (high - low) / 2
		currDistance := middle * (time - middle)
		if currDistance < target {
			low = middle + 1
		} else {
			high = middle
		}
	}

	return low
}

func UpperBoundBinarySearch(low, high, target int) int {
	time := high
	for low < high {
		middle := low + (high - low) / 2
		currDistance := middle * (time - middle)
		if currDistance < target {
			high = middle
		} else {
			low = middle + 1
		}
	}

	return high
}