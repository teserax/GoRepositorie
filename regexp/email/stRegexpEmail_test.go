package main

import "testing"

func TestEmailSearch(t *testing.T) {
	s := "dadadad steeles80@@@m34ail.ru ijsifsdfsd steeles80@mail.ru steeles80@m34ail..ru qeqeqw eqeqweq@sfsdf sadad42@3df@dfgdf.ouo stee/[les80@m34ail.ru"
	if col := emailSearch(s); col != "steeles80@mail.ru" {
		t.Errorf("Expected '%s', but got '%s'", "steeles80@mail.ru", col)
	}
}
