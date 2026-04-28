package piscine

func isValidBase2(base string) bool {
	if len(base) < 2 {
		return false
	}
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase2(base) {
		return 0
	}
	result := 0
	for _, r := range s {
		index := -1
		for i, br := range base {
			if r == br {
				index = i
				break
			}
		}
		if index == -1 {
			return 0
		}
		result = result*len(base) + index
	}
	return result
}
