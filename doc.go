/*
Package inquote translates runes into their representation inside a quoted string,

For example, rune value 10 (which is the "line feed" character) becomes \n

Also, for example, rune value 31 (which is the "unit separator" character) becomes \x1f

Sample Code

	
	// ...
	
	r1 := rune(11) // This is the Vertical Tab character. Which could also be written '\v'
	r2 := rune(84) // This is the Uppercase T character. Which could also be written 'T'
	r2 := rune(30) // This is the Row Separator character. Which could also be written '\x1e'
	
	p := []byte{}
	
	p = inquote.AppendRune(p, r1)
	// p == []byte{ rune(92), rune(118) } == []byte{'\\', 'v'}
	
	p = inquote.AppendRune(p, r2)
	// p == []byte{ rune(92), rune(118), rune(84) } == []byte{'\\', 'v', 'T'}
	
	p = inquote.AppendRune(p, r3)
	// p == []byte{ rune(92), rune(118), rune(84), rune(92), rune(120), rune(49), rune(101) } == []byte{'\\', 'v', 'T', '\\', 'x', '1', 'e'}


*/
package inquote
