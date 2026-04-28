package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	direction := 0

	// Find direction
	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])
		if cmp != 0 {
			direction = cmp
			break
		}
	}

	// All equal case
	if direction == 0 {
		return true
	}

	// Check consistency
	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])
		if cmp != 0 && (cmp > 0) != (direction > 0) {
			return false
		}
	}

	return true
}
