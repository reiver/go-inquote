package inquote


import (
	"bytes"
	"strconv"

	"testing"
)


func TestWriteRune(t *testing.T) {

	tests := []struct{
		Rune     rune
		Expected string
	}{
		{
			Rune:     rune(0), // NUL
			Expected: `\x00`,
		},
		{
			Rune:     rune(1), // SOH
			Expected: `\x01`,
		},
		{
			Rune:     rune(2), // STX
			Expected: `\x02`,
		},
		{
			Rune:     rune(3), // ETX
			Expected: `\x03`,
		},
		{
			Rune:     rune(4), // EOT
			Expected: `\x04`,
		},
		{
			Rune:     rune(5), // ENQ
			Expected: `\x05`,
		},
		{
			Rune:     rune(6), // ACK
			Expected: `\x06`,
		},
		{
			Rune:     rune(7), // BEL
			Expected: `\a`,
		},
		{
			Rune:     rune(8), // BS
			Expected: `\b`,
		},
		{
			Rune:     rune(9), //HT
			Expected: `\t`,
		},
		{
			Rune:     rune(10), // LF
			Expected: `\n`,
		},

		{
			Rune:     rune(11), // VT
			Expected: `\v`,
		},
		{
			Rune:     rune(12), // FF
			Expected: `\f`,
		},
		{
			Rune:     rune(13), // CR
			Expected: `\r`,
		},
		{
			Rune:     rune(14), // SO
			Expected: `\x0e`,
		},
		{
			Rune:     rune(15), // SI
			Expected: `\x0f`,
		},
		{
			Rune:     rune(16), // DLE
			Expected: `\x10`,
		},
		{
			Rune:     rune(17), // DC1
			Expected: `\x11`,
		},
		{
			Rune:     rune(18), // DC2
			Expected: `\x12`,
		},
		{
			Rune:     rune(19), // DC3
			Expected: `\x13`,
		},
		{
			Rune:     rune(20), // DC4
			Expected: `\x14`,
		},
		{
			Rune:     rune(21), // NAK
			Expected: `\x15`,
		},
		{
			Rune:     rune(22), // SYN
			Expected: `\x16`,
		},
		{
			Rune:     rune(23), // ETB
			Expected: `\x17`,
		},
		{
			Rune:     rune(24), // CAN
			Expected: `\x18`,
		},
		{
			Rune:     rune(25), // EM
			Expected: `\x19`,
		},
		{
			Rune:     rune(26), // SUB
			Expected: `\x1a`,
		},
		{
			Rune:     rune(27), // ESC
			Expected: `\x1b`,
		},
		{
			Rune:     rune(28), // FS
			Expected: `\x1c`,
		},
		{
			Rune:     rune(29), // GS
			Expected: `\x1d`,
		},
		{
			Rune:     rune(30), // RS
			Expected: `\x1e`,
		},
		{
			Rune:     rune(31), // US
			Expected: `\x1f`,
		},
		{
			Rune:     ' ',
			Expected: " ",
		},
		{
			Rune:     '!',
			Expected: "!",
		},
		{
			Rune:     '"',
			Expected: `\"`,
		},
		{
			Rune:     '#',
			Expected: "#",
		},
		{
			Rune:     '$',
			Expected: "$",
		},
		{
			Rune:     '%',
			Expected: "%",
		},
		{
			Rune:     '&',
			Expected: "&",
		},
		{
			Rune:     '\'',
			Expected: `'`,
		},
		{
			Rune:     '(',
			Expected: "(",
		},
		{
			Rune:     ')',
			Expected: ")",
		},
		{
			Rune:     '*',
			Expected: "*",
		},
		{
			Rune:     '+',
			Expected: "+",
		},
		{
			Rune:     ',',
			Expected: ",",
		},
		{
			Rune:     '-',
			Expected: "-",
		},
		{
			Rune:     '.',
			Expected: ".",
		},
		{
			Rune:     '/',
			Expected: "/",
		},
		{
			Rune:     '0',
			Expected: "0",
		},
		{
			Rune:     '1',
			Expected: "1",
		},
		{
			Rune:     '2',
			Expected: "2",
		},
		{
			Rune:     '3',
			Expected: "3",
		},
		{
			Rune:     '4',
			Expected: "4",
		},
		{
			Rune:     '5',
			Expected: "5",
		},
		{
			Rune:     '6',
			Expected: "6",
		},
		{
			Rune:     '7',
			Expected: "7",
		},
		{
			Rune:     '8',
			Expected: "8",
		},
		{
			Rune:     '9',
			Expected: "9",
		},
		{
			Rune:     ':',
			Expected: ":",
		},
		{
			Rune:     ';',
			Expected: ";",
		},
		{
			Rune:     '<',
			Expected: "<",
		},
		{
			Rune:     '=',
			Expected: "=",
		},
		{
			Rune:     '>',
			Expected: ">",
		},
		{
			Rune:     '?',
			Expected: "?",
		},
		{
			Rune:     '@',
			Expected: "@",
		},
		{
			Rune:     'A',
			Expected: "A",
		},
		{
			Rune:     'B',
			Expected: "B",
		},
		{
			Rune:     'C',
			Expected: "C",
		},
		{
			Rune:     'D',
			Expected: "D",
		},
		{
			Rune:     'E',
			Expected: "E",
		},
		{
			Rune:     'F',
			Expected: "F",
		},
		{
			Rune:     'G',
			Expected: "G",
		},
		{
			Rune:     'H',
			Expected: "H",
		},
		{
			Rune:     'I',
			Expected: "I",
		},
		{
			Rune:     'J',
			Expected: "J",
		},
		{
			Rune:     'K',
			Expected: "K",
		},
		{
			Rune:     'L',
			Expected: "L",
		},
		{
			Rune:     'M',
			Expected: "M",
		},
		{
			Rune:     'N',
			Expected: "N",
		},
		{
			Rune:     'O',
			Expected: "O",
		},
		{
			Rune:     'P',
			Expected: "P",
		},
		{
			Rune:     'Q',
			Expected: "Q",
		},
		{
			Rune:     'R',
			Expected: "R",
		},
		{
			Rune:     'S',
			Expected: "S",
		},
		{
			Rune:     'T',
			Expected: "T",
		},
		{
			Rune:     'U',
			Expected: "U",
		},
		{
			Rune:     'V',
			Expected: "V",
		},
		{
			Rune:     'W',
			Expected: "W",
		},
		{
			Rune:     'X',
			Expected: "X",
		},
		{
			Rune:     'Y',
			Expected: "Y",
		},
		{
			Rune:     'Z',
			Expected: "Z",
		},
		{
			Rune:     '[',
			Expected: "[",
		},
		{
			Rune:     '\\',
			Expected: `\\`,
		},
		{
			Rune:     ']',
			Expected: "]",
		},
		{
			Rune:     '^',
			Expected: "^",
		},
		{
			Rune:     '_',
			Expected: "_",
		},
		{
			Rune:     '`',
			Expected: "`",
		},
		{
			Rune:     'a',
			Expected: "a",
		},
		{
			Rune:     'b',
			Expected: "b",
		},
		{
			Rune:     'c',
			Expected: "c",
		},
		{
			Rune:     'd',
			Expected: "d",
		},
		{
			Rune:     'e',
			Expected: "e",
		},
		{
			Rune:     'f',
			Expected: "f",
		},
		{
			Rune:     'g',
			Expected: "g",
		},
		{
			Rune:     'h',
			Expected: "h",
		},
		{
			Rune:     'i',
			Expected: "i",
		},
		{
			Rune:     'j',
			Expected: "j",
		},
		{
			Rune:     'k',
			Expected: "k",
		},
		{
			Rune:     'l',
			Expected: "l",
		},
		{
			Rune:     'm',
			Expected: "m",
		},
		{
			Rune:     'n',
			Expected: "n",
		},
		{
			Rune:     'o',
			Expected: "o",
		},
		{
			Rune:     'p',
			Expected: "p",
		},
		{
			Rune:     'q',
			Expected: "q",
		},
		{
			Rune:     'r',
			Expected: "r",
		},
		{
			Rune:     's',
			Expected: "s",
		},
		{
			Rune:     't',
			Expected: "t",
		},
		{
			Rune:     'u',
			Expected: "u",
		},
		{
			Rune:     'v',
			Expected: "v",
		},
		{
			Rune:     'w',
			Expected: "w",
		},
		{
			Rune:     'x',
			Expected: "x",
		},
		{
			Rune:     'y',
			Expected: "y",
		},
		{
			Rune:     'z',
			Expected: "z",
		},

		{
			Rune:     '{',
			Expected: "{",
		},
		{
			Rune:     '|',
			Expected: "|",
		},
		{
			Rune:     '}',
			Expected: "}",
		},
		{
			Rune:     '~',
			Expected: "~",
		},
		{
			Rune:     rune(127),
			Expected: `\x7f`,
		},



		{
			Rune:     '✪', // CIRCLED WHITE STAR
			Expected: "\u272A",
		},



		{
			Rune:     rune(8232), // LINE SEPARATOR
			Expected: `\u2028`,
		},



		{
			Rune:     '🙂', // SLIGHTLY SMILING FACE
			Expected: "\U0001F642",
		},
	}


	for testNumber, test := range tests {
		var buffer bytes.Buffer

		_, err := WriteRune(&buffer, test.Rune)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		{
			expectedString := test.Expected
			actualString   := buffer.String()

			if expected, actual := len(expectedString), len(actualString); expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				continue
			}

			if expected, actual := expectedString, actualString;  expected != actual {
				t.Errorf("For test #%d, expected ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
				continue
			}
		}

		// For some reason Golang encodes character 127 in \u format. So skipping that.
		if rune(127) != test.Rune {
			quoted := strconv.Quote( string(test.Rune) )
			should := quoted[1:len(quoted)-1]

			if expected, actual := should, buffer.String(); expected != actual {
				t.Errorf("For test #%d, expected that it should be ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
				continue
			}
		}
	}
}
