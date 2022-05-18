package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	n := []rune(input)
	m := make([]rune, len(n))
	for i, s := 0, len(n)-1; i != len(n); i, s = i+1, s-1 {
		m[i] = n[s]
	}

	output = string(m)

	return output
}
