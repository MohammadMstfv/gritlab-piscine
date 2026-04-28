package piscine

func IsLower(s string) bool {
	for _, r2 := range s {
		if r2 > 122 || r2 < 97 {
			return false
		}
	}
	return true
}
