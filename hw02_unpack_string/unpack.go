package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (result string, err error) {
	if _, err := strconv.Atoi(s); err == nil {
		return result, ErrInvalidString
	}
	var prev rune
	var escaped, next bool
	var b strings.Builder
	for i, char := range s {
		if (i == 0 && unicode.IsDigit(char)) || (next && unicode.IsDigit(char)) {
			return result, ErrInvalidString
		}
		if unicode.IsDigit(char) && !escaped {
			next = true
			var m int
			if char == '0' {
				c := b.String()
				c = c[:len(c)-1]
				b.Reset()
				b.WriteString(c)
			} else {
				m = int(char - '0')
				result = strings.Repeat(string(prev), m-1)
				b.WriteString(result)
			}
		} else {
			next = false
			escaped = string(char) == "\\" && string(prev) != "\\"
			if !escaped {
				b.WriteRune(char)
			}
			prev = char
		}
	}

	return b.String(), err
}
