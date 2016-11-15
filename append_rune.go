package inquote


import (
	"unicode"
	"unicode/utf8"
)


// AppendRune appends the string-literal-encoding of a rune to the []byte `p`.
//
// Many runes stay as the are.
//
// For example, rune(55), rune(65), and rune(100) would stay as is, since they
// are the runes '7', 'A' and 'd' respectively.
//
// However, some runes change.
//
// For example, rune(13), rune(21), rune(27), and rune(92) would change, since they are
// the rune runes '\r', '\x15', '\x1b', and '\\' respectively.
func AppendRune(p []byte, r rune) []byte {

	if r > utf8.MaxRune {
		return append(p, '\\', 'u', 'f', 'f', 'f', 'd')
	}

	if '\\' == r {
		return append(p, `\\`...)
	}

	if '"' == r {
		return append(p, `\"`...)
	}

	switch r {
	case '\a':
		return append(p, `\a`...)
	case '\b':
		return append(p, `\b`...)
	case '\f':
		return append(p, `\f`...)
	case '\n':
		return append(p, `\n`...)
	case '\r':
		return append(p, `\r`...)
	case '\t':
		return append(p, `\t`...)
	case '\v':
		return append(p, `\v`...)
	}

	if !unicode.IsPrint(r) && r < utf8.RuneSelf {
		c0 := byte((r     ) % 16)
		c1 := byte((r >> 4) % 16)

		if 0 <= c0 && c0 <= 9 {
			c0 = byte('0') + c0
		} else {
			c0 = byte('a') + c0 - 10
		}

		if 0 <= c1 && c1 <= 9 {
			c1 = byte('0') + c1
		} else {
			c1 = byte('a') + c1 - 10
		}

		return append(p, '\\', 'x', c1, c0)
	}

	if !unicode.IsPrint(r) && 0x10000 > r {
		c0 := byte((r      ) % 16)
		c1 := byte((r >>  4) % 16)
		c2 := byte((r >>  8) % 16)
		c3 := byte((r >> 12) % 16)

		if 0 <= c0 && c0 <= 9 {
			c0 = byte('0') + c0
		} else {
			c0 = byte('a') + c0 - 10
		}

		if 0 <= c1 && c1 <= 9 {
			c1 = byte('0') + c1
		} else {
			c1 = byte('a') + c1 - 10
		}

		if 0 <= c2 && c2 <= 9 {
			c2 = byte('0') + c2
		} else {
			c2 = byte('a') + c2 - 10
		}

		if 0 <= c3 && c3 <= 9 {
			c3 = byte('0') + c3
		} else {
			c3 = byte('a') + c3 - 10
		}

		return append(p, '\\', 'u', c3, c2, c1, c0)
	}

	if !unicode.IsPrint(r) && 0x10000 <= r {
		c0 := byte((r      ) % 16)
		c1 := byte((r >>  4) % 16)
		c2 := byte((r >>  8) % 16)
		c3 := byte((r >> 12) % 16)
		c4 := byte((r >> 16) % 16)
		c5 := byte((r >> 20) % 16)
		c6 := byte((r >> 24) % 16)
		c7 := byte((r >> 28) % 16)

		if 0 <= c0 && c0 <= 9 {
			c0 = byte('0') + c0
		} else {
			c0 = byte('a') + c0 - 10
		}

		if 0 <= c1 && c1 <= 9 {
			c1 = byte('0') + c1
		} else {
			c1 = byte('a') + c1 - 10
		}

		if 0 <= c2 && c2 <= 9 {
			c2 = byte('0') + c2
		} else {
			c2 = byte('a') + c2 - 10
		}

		if 0 <= c3 && c3 <= 9 {
			c3 = byte('0') + c3
		} else {
			c3 = byte('a') + c3 - 10
		}

		if 0 <= c4 && c4 <= 9 {
			c4 = byte('0') + c4
		} else {
			c4 = byte('a') + c4 - 10
		}

		if 0 <= c5 && c5 <= 9 {
			c5 = byte('0') + c5
		} else {
			c5 = byte('a') + c5 - 10
		}

		if 0 <= c6 && c6 <= 9 {
			c6 = byte('0') + c6
		} else {
			c6 = byte('a') + c6 - 10
		}

		if 0 <= c7 && c7 <= 9 {
			c7 = byte('0') + c7
		} else {
			c7 = byte('a') + c7 - 10
		}

		return append(p, '\\', 'U', c7, c6, c5, c4, c3, c2, c1, c0)
	}

	var buffer [utf8.UTFMax]byte
	n := utf8.EncodeRune(buffer[:], r)
	return append(p, buffer[:n]...)
}
