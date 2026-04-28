package piscine

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	resulo := 1
	for i := nb; i >= 1; i-- {
		resulo = resulo * i
	}
	return resulo
}
