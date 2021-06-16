package runeutil

import "unicode"

func IsBlankRune(r rune) bool {
	return isSpaceRune(r) || r == '\ufeff' || r == '\u202a'
}

func isSpaceRune(r rune) bool {
	return unicode.IsSpace(r)
}
