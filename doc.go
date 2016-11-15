/*
Package inquote translates runes into their representation inside a quoted string,

For example, rune value 10 (which is the "line feed" character) becomes \n

Also, for example, rune value 31 (which is the "unit separator" character) becomes \x1f

Sample Code Using AppendRune

	p := []byte{}
	
	// ...
	
	r := rune(11) // This is the Vertical Tab character. Which could also be written '\v'
	
	p = inquote.AppendRune(p, r)
	
	// p == []byte{ rune(92), rune(118) } == []byte{'\\', 'v'}

*/
package inquote
