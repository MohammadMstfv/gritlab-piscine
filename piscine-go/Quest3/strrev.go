package piscine

func StrRev(s string) string {
	runes := []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		runes = append(runes, rune(s[i]))
	}
	return string(runes)
}
