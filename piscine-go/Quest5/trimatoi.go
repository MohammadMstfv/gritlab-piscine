package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := false
	foundDigit := false

	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
			foundDigit = true
		} else if r == '-' && !foundDigit {
			negative = true
		}
	}

	if negative {
		result = -result
	}

	return result
}
