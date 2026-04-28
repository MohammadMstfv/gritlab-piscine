package main

import "os"

func writeStr(s string) {
	os.Stdout.Write([]byte(s))
}

func writeInt(n int64) {
	if n == 0 {
		writeStr("0\n")
		return
	}

	if n < 0 {
		writeStr("-")
		n = -n
	}

	var digits [20]byte
	i := 0

	for n > 0 {
		digits[i] = byte(n%10) + '0'
		n /= 10
		i++
	}

	for i > 0 {
		i--
		os.Stdout.Write([]byte{digits[i]})
	}
	os.Stdout.Write([]byte{'\n'})
}

func atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}

	sign := int64(1)
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	var res int64 = 0

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int64(s[i] - '0')

		// overflow check
		if res > (9223372036854775807-digit)/10 {
			return 0, false
		}

		res = res*10 + digit
	}

	return res * sign, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	b, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	op := os.Args[2]
	if len(op) != 1 {
		return
	}

	switch op[0] {
	case '+':
		res := a + b
		if (b > 0 && res < a) || (b < 0 && res > a) {
			return
		}
		writeInt(res)

	case '-':
		res := a - b
		if (b < 0 && res < a) || (b > 0 && res > a) {
			return
		}
		writeInt(res)

	case '*':
		res := a * b
		if a != 0 && res/a != b {
			return
		}
		writeInt(res)

	case '/':
		if b == 0 {
			writeStr("No division by 0\n")
			return
		}
		writeInt(a / b)

	case '%':
		if b == 0 {
			writeStr("No modulo by 0\n")
			return
		}
		writeInt(a % b)

	default:
		return
	}
}
