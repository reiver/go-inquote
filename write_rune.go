package inquote


import (
	"io"
)


// WriteRune writes the string-literal-encoding of a rune to the io.Writer `w`.
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
func WriteRune(w io.Writer, r rune) (int, error) {
	if nil == w {
		return 0, errNilWriter
	}

	var buffer [10]byte // Longest for is \U12345678

	p := AppendRune(buffer[:0], r)

	return w.Write(p)
}
