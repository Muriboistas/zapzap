package stringx

// ToArgs recieve a string and return args
func ToArgs(text string) []string {
	var args []string
	for len(text) > 0 {
		if text[0] == ' ' || text[0] == '\t' {
			text = text[1:]
			continue
		}
		var arg []byte
		arg, text = readNextArg(text)
		args = append(args, string(arg))
	}
	return args
}

func readNextArg(text string) (arg []byte, rest string) {
	var b []byte
	var inquote bool
	var nslash int
	for ; len(text) > 0; text = text[1:] {
		c := text[0]
		switch c {
		case ' ', '\t':
			if !inquote {
				return appendBSBytes(b, nslash), text[1:]
			}
		case '"':
			b = appendBSBytes(b, nslash/2)
			if nslash%2 == 0 {
				// use "Prior to 2008" rule from
				// http://daviddeley.com/autohotkey/parameters/parameters.htm
				// section 5.2 to deal with double double quotes
				if inquote && len(text) > 1 && text[1] == '"' {
					b = append(b, c)
					text = text[1:]
				}
				inquote = !inquote
			} else {
				b = append(b, c)
			}
			nslash = 0
			continue
		case '\\':
			nslash++
			continue
		}
		b = appendBSBytes(b, nslash)
		nslash = 0
		b = append(b, c)
	}
	return appendBSBytes(b, nslash), ""
}

// appendBSBytes appends n '\\' bytes to b and returns the resulting slice.
func appendBSBytes(b []byte, n int) []byte {
	for ; n > 0; n-- {
		b = append(b, '\\')
	}
	return b
}
