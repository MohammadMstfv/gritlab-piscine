package piscine

func ToLower(s string) string {
	result := ""
	for _, r2 := range s {
		if r2 >= 65 && r2 <= 90 {
			r2 += 32
		}
		result += string(r2)
	}
	return result
}
