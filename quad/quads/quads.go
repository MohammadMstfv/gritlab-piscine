package quads

import "strings"

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (j == 0 && i == 0) || (j == x-1 && i == 0) || (j == 0 && i == y-1) || (j == x-1 && i == y-1) {
				sb.WriteString("o")
			} else if (j != 0 && i == 0) || (j != 0 && i == y-1) {
				sb.WriteString("-")
			} else if j == x-1 || j == 0 {
				sb.WriteString("|")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if j == 0 && i == 0 {
				sb.WriteString("/")
			} else if j == x-1 && i == 0 {
				sb.WriteString("\\")
			} else if j == 0 && i == y-1 {
				sb.WriteString("\\")
			} else if j == x-1 && i == y-1 {
				sb.WriteString("/")
			} else if (j != 0 && i == 0) || (j != 0 && i == y-1) || (j == x-1 || j == 0) {
				sb.WriteString("*")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func QuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if j == 0 && i == 0 {
				sb.WriteString("A")
			} else if j == x-1 && i == 0 {
				sb.WriteString("A")
			} else if j == 0 && i == y-1 {
				sb.WriteString("C")
			} else if j == x-1 && i == y-1 {
				sb.WriteString("C")
			} else if (j != 0 && i == 0) || (j != 0 && i == y-1) || (j == x-1 || j == 0) {
				sb.WriteString("B")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if j == 0 && i == 0 {
				sb.WriteString("A")
			} else if j == x-1 && i == 0 {
				sb.WriteString("C")
			} else if j == 0 && i == y-1 {
				sb.WriteString("A")
			} else if j == x-1 && i == y-1 {
				sb.WriteString("C")
			} else if (j != 0 && i == 0) || (j != 0 && i == y-1) || (j == x-1 || j == 0) {
				sb.WriteString("B")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func QuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if j == 0 && i == 0 {
				sb.WriteString("A")
			} else if j == x-1 && i == 0 {
				sb.WriteString("C")
			} else if j == 0 && i == y-1 {
				sb.WriteString("C")
			} else if j == x-1 && i == y-1 {
				sb.WriteString("A")
			} else if (j != 0 && i == 0) || (j != 0 && i == y-1) || (j == x-1 || j == 0) {
				sb.WriteString("B")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
