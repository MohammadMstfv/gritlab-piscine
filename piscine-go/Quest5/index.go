// My Code
package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	rs := []rune(s)
	rfind := []rune(toFind)
	i := 0
	j := 0
outer:
	for i < len(rs) {
		if rs[i] == rfind[0] {
			goto subStrCheck
		}
		i++
	}
	goto End
subStrCheck:
	j = 0
	for j < len(rfind) {
		if rs[i] != rfind[j] {
			i = i - j + 1
			j = 0
			goto outer
		}
		j++
		i++
	}
	return i - j
End:
	return -1
}

//

/* Claude Code
package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
*/

/* ChatGPT Code, rune safe solution
package piscine

func Index(s, toFind string) int {
	rs := []rune(s)
	rfind := []rune(toFind)

	for i := 0; i <= len(rs)-len(rfind); i++ {
		match := true
		for j := range rfind {
			if rs[i+j] != rfind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
*/
