package inquote


import (
	"io"
)


func WriteRune(w io.Writer, r rune) (int, error) {
	if nil == w {
		return 0, errNilWriter
	}

	var buffer [10]byte // Longest for is \U12345678

	p := AppendRune(buffer[:0], r)

	return w.Write(p)
}
