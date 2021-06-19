package leftpad

import (
	"strings"
	"unicode/utf8"
)

/*
	default_char e' visibile solo all'interno del package, perche' il nome della variabile comincia con la lettera minuscola
*/
var default_char = ' '

func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
