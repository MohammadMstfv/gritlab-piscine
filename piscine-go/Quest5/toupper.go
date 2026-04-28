package piscine

func ToUpper(s string) string {
	result := ""
	for _, r2 := range s {
		if r2 >= 97 && r2 <= 122 {
			r2 -= 32
		}
		result += string(r2)
	}
	return result
}
