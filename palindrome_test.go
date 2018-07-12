package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	s1 := "evve"
	s1Is := isPalindrome(s1)

	if !s1Is {
		t.Errorf("Expected evve to be a palindrome")
	}

	s2 := "adf"
	s2Is := isPalindrome(s2)

	if s2Is {
		t.Errorf("Expected adf to not be a palindrome")
	}

	s3 := "aerea"
	s3Is := isPalindrome(s3)

	if !s3Is {
		t.Errorf("Expected aerea to be a palindrome")
	}
}