package piscine

func IsUpper(s string) bool {
	for _, r2 := range s {
		if r2 > 90 || r2 < 65 {
			return false
		}
	}
	return true
}
