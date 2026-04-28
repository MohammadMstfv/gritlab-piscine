package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert from baseFrom to decimal
	decimal := 0
	baseFromLen := len(baseFrom)

	for _, r := range nbr {
		digit := 0
		for i, b := range baseFrom {
			if b == r {
				digit = i
				break
			}
		}
		decimal = decimal*baseFromLen + digit
	}

	// Convert from decimal to baseTo
	if decimal == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	result := ""

	for decimal > 0 {
		remainder := decimal % baseToLen
		result = string(baseTo[remainder]) + result
		decimal /= baseToLen
	}

	return result
}
