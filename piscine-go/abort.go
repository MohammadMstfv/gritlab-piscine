package piscine

func Abort(a, b, c, d, e int) int {
	tempSlice := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4-i; j++ { // sort
			if tempSlice[j] > tempSlice[j+1] {
				tempIndex := tempSlice[j]
				tempSlice[j] = tempSlice[j+1]
				tempSlice[j+1] = tempIndex
			}
		}
	}
	return tempSlice[2]
}
