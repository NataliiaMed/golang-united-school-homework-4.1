package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	l := len(input)
	m := make([]rune, l)
	for _, s := range(input) {
		l--
		m[l] = s
	}output = string(m)
	
return output
}
	

	
