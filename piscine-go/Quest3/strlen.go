package piscine

func StrLen(s string) int {
	runecounter := 0
	for range s {
		runecounter++
	}
	return runecounter
}
