package main

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		chr := word[i]
		opp := word[len(word) - i - 1]

		if chr != opp {
			return false
		}
	}
	return true
}
