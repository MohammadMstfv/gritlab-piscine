package piscine

func SplitWhiteSpaces(s string) []string {
	mystr := make([]string, 0)
	currentWord := ""
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if currentWord != "" {
				mystr = append(mystr, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(r)
		}
	}
	if currentWord != "" {
		mystr = append(mystr, currentWord)
	}
	return mystr
}
