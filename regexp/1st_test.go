package main

import "testing"

func TestPriceTranslation(t *testing.T) {
	if col := priceTranslation("+&nbsp;EUR 18,00&nbsp;di&nbsp;spedizione"); col != "EUR 18,00" {
		t.Errorf("Expected '%s', but got '%s'", "EUR 18,00", col)
	}
}
