package piscine

func LastRune(s string) rune {
	nb := StrLen(s)
	for index, letter := range s {
		if index == nb-1 {
			return letter
		}
	}
	return -1
}
