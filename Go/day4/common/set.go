package common

func MakeSet(slice []int) (set map[int]bool) {
	set = make(map[int]bool)
	for _, num := range slice {
		set[num] = true
	}

	return set
}