package piscine

func IsPrintable(s string) bool {
	for _, r2 := range s {
		if r2 < 32 || r2 > 126 {
			return false
		}
	}
	return true
}
