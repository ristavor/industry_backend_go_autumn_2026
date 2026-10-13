package main

func rotateRunes(s string, shift int) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	i := shift % len(runes)
	if i < 0 {
		i = len(runes) + i
	}
	return string(append(runes[i:], runes[:i]...))
}
