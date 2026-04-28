package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		res := []int(nil)
		return res
	}
	size := max - min
	res := make([]int, size)
	for i := min; i < max; i++ {
		res[i-min] = i
	}
	return res
}
