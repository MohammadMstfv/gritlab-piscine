package piscine

func IsAlpha(s string) bool {
	for _, r2 := range s {
		if (r2 < 48 || r2 > 57) && (r2 < 65 || r2 > 90) && (r2 < 97 || r2 > 122) {
			return false
		}
	}
	return true
}
