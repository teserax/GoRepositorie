package main

import "testing"

type Sum struct {
	s, expected string
}

var summa = []Sum{
	Sum{"EUR 1.203,09", "EUR 1.203,09"},
	Sum{"$30.65", "$30.65"},
	Sum{"$30.65 ($13.74 / Pound)", "$30.65"},
	//Sum{"$1,000,010.46", "$1,000,010.46"},
	Sum{"+&nbsp;EUR 18,00&nbsp;di&nbsp;spedizione", "EUR 18,00"},
}

func TestCurrency(t *testing.T) {
	for i, test := range summa {
		if col := currency(test.s); col != test.expected {
			t.Errorf("%d.Expected '%s', but got '%s'", i, test.expected, col)
		}
	}

}
