package piscine

func ConcatParams(args []string) string {
	mid := make([]string, len(args)*2)
	result := ""
	for i, j := 0, 0; i < len(args)*2; i += 2 {
		mid[i] = args[j]
		j++
	}
	for i := 1; i < len(args)*2-1; i += 2 {
		mid[i] = "\n"
	}
	for _, r := range mid {
		result += r
	}

	return string(result)
}
