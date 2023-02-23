package recursion

func Reverse(input string) string {
	if input == "" {
		return ""
	}
	return Reverse(input[1:]) + string(input[0])
}
