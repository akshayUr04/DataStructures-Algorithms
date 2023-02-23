package recursion

func Palindrome(input string) bool {

	if len(input) == 0 || len(input) == 1 {
		return true
	}

	if string(input[0]) == string(input[len(input)-1]) {
		return Palindrome(input[1 : len(input)-1])
	}
	return false
}
