package main

import (
	"testing"
)

func TestPolidrome(t *testing.T) {
	s := []string{
		"ттопотт",
		"aabbaa",
		"fire",
		"wo",
	}
	for _, testStr := range s {
		tes := polindrome(testStr)
		if tes != true {
			t.Errorf("%s is not polindrom", testStr)
		}
	}
}
