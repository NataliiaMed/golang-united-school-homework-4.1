package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	m := make([]rune, len(input))
	for i, s := 0, len(input)-1; i < s; i, s = i+1, s-1 {

		m[i], m[s] = m[s], m[i]
	}

	output = string(m)

	return output
}
