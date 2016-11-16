package inquote


import (
	"testing"
)


func TestDecodeRune(t *testing.T) {

	tests := []struct{
		ExpectedEncodedLength int
		Expected              rune
		String                string
	}{
		{
			Expected: rune(0), // NUL
			String: `\x00`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(1), // SOH
			String: `\x01`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(2), // STX
			String: `\x02`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(3), // ETX
			String: `\x03`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(4), // EOT
			String: `\x04`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(5), // ENQ
			String: `\x05`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(6), // ACK
			String: `\x06`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(7), // BEL
			String: `\a`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: rune(8), // BS
			String: `\b`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: rune(9), //HT
			String: `\t`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: rune(10), // LF
			String: `\n`,
			ExpectedEncodedLength: 2,
		},

		{
			Expected: rune(11), // VT
			String: `\v`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: rune(12), // FF
			String: `\f`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: rune(13), // CR
			String: `\r`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: rune(14), // SO
			String: `\x0e`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(15), // SI
			String: `\x0f`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(16), // DLE
			String: `\x10`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(17), // DC1
			String: `\x11`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(18), // DC2
			String: `\x12`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(19), // DC3
			String: `\x13`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(20), // DC4
			String: `\x14`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(21), // NAK
			String: `\x15`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(22), // SYN
			String: `\x16`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(23), // ETB
			String: `\x17`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(24), // CAN
			String: `\x18`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(25), // EM
			String: `\x19`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(26), // SUB
			String: `\x1a`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(27), // ESC
			String: `\x1b`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(28), // FS
			String: `\x1c`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(29), // GS
			String: `\x1d`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(30), // RS
			String: `\x1e`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(31), // US
			String: `\x1f`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: ' ',
			String:   " ",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '!',
			String:   "!",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '"',
			String:  `\"`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: '#',
			String:   "#",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '$',
			String:   "$",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '%',
			String:   "%",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '&',
			String:   "&",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '\'',
			String:    `'`,
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '(',
			String:   "(",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: ')',
			String:   ")",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '*',
			String:   "*",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '+',
			String:   "+",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: ',',
			String:   ",",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '-',
			String:   "-",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '.',
			String:   ".",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '/',
			String:   "/",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '0',
			String:   "0",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '1',
			String:   "1",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '2',
			String:   "2",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '3',
			String:   "3",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '4',
			String:   "4",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '5',
			String:   "5",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '6',
			String:   "6",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '7',
			String:   "7",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '8',
			String:   "8",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '9',
			String:   "9",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: ':',
			String:   ":",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: ';',
			String:   ";",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '<',
			String:   "<",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '=',
			String:   "=",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '>',
			String:   ">",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '?',
			String:   "?",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '@',
			String:   "@",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'A',
			String:   "A",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'B',
			String:   "B",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'C',
			String:   "C",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'D',
			String:   "D",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'E',
			String:   "E",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'F',
			String:   "F",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'G',
			String:   "G",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'H',
			String:   "H",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'I',
			String:   "I",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'J',
			String:   "J",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'K',
			String:   "K",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'L',
			String:   "L",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'M',
			String:   "M",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'N',
			String:   "N",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'O',
			String:   "O",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'P',
			String:   "P",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'Q',
			String:   "Q",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'R',
			String:   "R",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'S',
			String:   "S",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'T',
			String:   "T",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'U',
			String:   "U",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'V',
			String:   "V",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'W',
			String:   "W",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'X',
			String:   "X",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'Y',
			String:   "Y",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'Z',
			String:   "Z",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '[',
			String:   "[",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '\\',
			String:   `\\`,
			ExpectedEncodedLength: 2,
		},
		{
			Expected: ']',
			String:   "]",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '^',
			String:   "^",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '_',
			String:   "_",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '`',
			String:   "`",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'a',
			String:   "a",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'b',
			String:   "b",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'c',
			String:   "c",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'd',
			String:   "d",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'e',
			String:   "e",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'f',
			String:   "f",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'g',
			String:   "g",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'h',
			String:   "h",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'i',
			String:   "i",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'j',
			String:   "j",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'k',
			String:   "k",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'l',
			String:   "l",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'm',
			String:   "m",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'n',
			String:   "n",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'o',
			String:   "o",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'p',
			String:   "p",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'q',
			String:   "q",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'r',
			String:   "r",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 's',
			String:   "s",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 't',
			String:   "t",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'u',
			String:   "u",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'v',
			String:   "v",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'w',
			String:   "w",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'x',
			String:   "x",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'y',
			String:   "y",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: 'z',
			String:   "z",
			ExpectedEncodedLength: 1,
		},

		{
			Expected: '{',
			String:   "{",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '|',
			String:   "|",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '}',
			String:   "}",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: '~',
			String:   "~",
			ExpectedEncodedLength: 1,
		},
		{
			Expected: rune(127),
			String: `\x7f`,
			ExpectedEncodedLength: 4,
		},



		{
			Expected: '✪', // CIRCLED WHITE STAR
			String:   `\u272a`,
			ExpectedEncodedLength: 6,
		},



		{
			Expected: rune(8232), // LINE SEPARATOR
			String: `\u2028`,
			ExpectedEncodedLength: 6,
		},



		{
			Expected: '🙂', // SLIGHTLY SMILING FACE
			String:   `\U0001f642`,
			ExpectedEncodedLength: 10,
		},






		{
			Expected: rune(14), // SO
			String: `\x0E`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(15), // SI
			String: `\x0F`,
			ExpectedEncodedLength: 4,
		},



		{
			Expected: rune(26), // SUB
			String: `\x1A`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(27), // ESC
			String: `\x1B`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(28), // FS
			String: `\x1C`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(29), // GS
			String: `\x1D`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(30), // RS
			String: `\x1E`,
			ExpectedEncodedLength: 4,
		},
		{
			Expected: rune(31), // US
			String: `\x1F`,
			ExpectedEncodedLength: 4,
		},


		{
			Expected: rune(127),
			String: `\x7F`,
			ExpectedEncodedLength: 4,
		},




		{
			Expected: '✪', // CIRCLED WHITE STAR
			String:   `\u272A`,
			ExpectedEncodedLength: 6,
		},



		{
			Expected: '🙂', // SLIGHTLY SMILING FACE
			String:   `\U0001F642`,
			ExpectedEncodedLength: 10,
		},
	}


	for testNumber, test := range tests {

		r, n, err := DecodeRune( []byte(test.String) )
		if nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if expected, actual := test.ExpectedEncodedLength, n; expected != actual {
			t.Errorf("For test #%d, expected the encoded length to be %d, but actually was %d. (String: %q)", testNumber, expected, actual, test.String)
			continue
		}
		if expected, actual := test.Expected, r; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, string(expected), string(actual))
			continue
		}


//@TODO: Concatenate stuff to end of string.

	}
}
