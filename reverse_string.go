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
	for i, s := 0, len(m)-1; i < s; i, s = i+1, s-1 {
		m[i], m[s] = m[s], m[i]
	}
	output = string(m)

	return output
}
