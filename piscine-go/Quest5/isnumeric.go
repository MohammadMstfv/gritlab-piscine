package piscine

func IsNumeric(s string) bool {
	for _, r2 := range s {
		if r2 < 48 || r2 > 57 {
			return false
		}
	}
	return true
}
