package inquote


import (
	"errors"
	"io"
)


var (
	errBadRequest = errors.New("Bad Request")
)


// DecodeRune decodes a string-literal-encoding of a rune from the beginning of `p`.
//
// It returns the value of the return, and the number of bytes the rune occupied.
//
// So, for example, one code do decoding with code such as:
//
//	var p []byte
//	p = ...
//	
//	for 0 < len(p) {
//		r, n, err := inquote.DecodeRune(p)
//		if nil != err {
//			if io.EOF == err {
//				break
//			}
//			return err
//		}
//		
//		fmt.Printf("RUNE: %d\n", r)
//		
//		p = p[n:]
//	}
//
func DecodeRune(p []byte) (rune, int, error) {

	lenp := len(p)

	if lenp < 1 {
		return 0, 0, io.EOF
	}


	r0 := rune(p[0])
	if '\\' == r0 {
		if lenp < 2 {
			return 0, 0, errBadRequest
		}

		r1 := rune(p[1])
		switch r1 {
		case '"':
			return '"',  2, nil
		case '\\':
			return '\\', 2, nil
		case 'a':
			return '\a', 2, nil
		case 'b':
			return '\b', 2, nil
		case 'f':
			return '\f', 2, nil
		case 'n':
			return '\n', 2, nil
		case 'r':
			return '\r', 2, nil
		case 't':
			return '\t', 2, nil
		case 'v':
			return '\v', 2, nil
		case 'u':
			if lenp < 6 {
				return 0, 0, errBadRequest
			}

			x := rune(0)

			for i:=2; i <= 5; i++ {

				x <<= 4

				r := rune(p[i])

				if '0' <= r && r <= '9'  {
					x += r - '0'
				} else if 'a' <= r && r <= 'f' {
					x += r - 'a' + 10
				} else if 'A' <= r && r <= 'F' {
					x += r - 'A' + 10
				} else {
					return 0, 0, errBadRequest
				}
			}

			return x, 6, nil
		case 'U':
			if lenp < 6 {
				return 0, 0, errBadRequest
			}

			x := rune(0)

			for i:=2; i <= 9; i++ {

				x <<= 4

				r := rune(p[i])

				if '0' <= r && r <= '9'  {
					x += r - '0'
				} else if 'a' <= r && r <= 'f' {
					x += r - 'a' + 10
				} else if 'A' <= r && r <= 'F' {
					x += r - 'A' + 10
				} else {
					return 0, 0, errBadRequest
				}
			}

			return x, 10, nil
		case 'x':
			if lenp < 4 {
				return 0, 0, errBadRequest
			}

			x := rune(0)

			for i:=2; i <= 3; i++ {

				x <<= 4

				r := rune(p[i])

				if '0' <= r && r <= '9'  {
					x += r - '0'
				} else if 'a' <= r && r <= 'f' {
					x += r - 'a' + 10
				} else if 'A' <= r && r <= 'F' {
					x += r - 'A' + 10
				} else {
					return 0, 0, errBadRequest
				}
			}

			return x, 4, nil
		default:
			// Nothing here.
		}
	}


	return r0, 1, nil
}
