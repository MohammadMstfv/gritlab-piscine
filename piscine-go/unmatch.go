package piscine

func Unmatch(a []int) int {
	if len(a) == 0 {
		return -1
	}

	count := make(map[int]int)

	for _, v := range a {
		count[v]++
	}

	for _, v := range a {
		if count[v]%2 != 0 {
			return v
		}
	}

	return -1
}
