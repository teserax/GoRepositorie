package main

import "testing"

type Test struct {
	s, expected string
}

var denominationTest = []Test{
	Test{"EUR 1.203,09", "EUR 1.203,09"},
	Test{"$30.65", "$30.65"},
	Test{"$30.65 ($13.74 / Pound)", "$30.65"},
	//Test{"$1,000,010.46", "$1,000,010.46"},
	Test{"+&nbsp;EUR 18,00&nbsp;di&nbsp;spedizione", "EUR 18,00"},
	Test{"MP3", ""},
}

func TestCurrency(t *testing.T) {
	for i, test := range denominationTest {
		if col := currency(test.s); col != test.expected {
			t.Errorf("%d.Expected '%s', but got '%s'", i, test.expected, col)
		}
	}

}
