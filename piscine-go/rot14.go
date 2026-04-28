package piscine

func Rot14(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 14
			if r > 90 { // after Z
				r -= 26
			}
		} else if r >= 'a' && r <= 'z' {
			r += 14
			if r > 122 { // after z
				r -= 26
			}
		}
		result += string(r)
	}
	return result
}
