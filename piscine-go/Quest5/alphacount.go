package piscine

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) { // a-z or A-Z
			count++
		}
	}
	return count
}
