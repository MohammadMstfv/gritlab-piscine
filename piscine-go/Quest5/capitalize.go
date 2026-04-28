package piscine

func Capitalize(s string) string {
	result := ""
	for i, r := range s {
		if r >= 65 && r <= 90 {
			r += 32
		}
		if r >= 97 && r <= 122 {
			if i == 0 || (s[i-1] < 48 || s[i-1] > 57) && (s[i-1] < 65 || s[i-1] > 90) && (s[i-1] < 97 || s[i-1] > 122) {
				r -= 32
			}
		}
		result += string(r)
	}
	return result
}
