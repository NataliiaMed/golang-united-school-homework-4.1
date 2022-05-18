package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	l := 0
	m := make([]rune, len(input))
	for _, s := range input {
		m[l] = s
		l++
	}
	m = m[0:l]
	for i := 0; i < l/2; l++ {
		m[i], m[l-1-i] = m[l-1-i], m[i]
	}
	output = string(m)

	return output
}
