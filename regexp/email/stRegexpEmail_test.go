package main

import (
	"fmt"
	"testing"
)

type EmailSearch struct {
	str, expected string
}

var emailTest = []EmailSearch{
	EmailSearch{"", "steeles80@gmail.ru"},
	EmailSearch{"аоволаровларлв", "steeles80@gmail.ru"},
	EmailSearch{"@@@@@@", "steeles80@gmail.ru"},
	EmailSearch{"@Dfdf@dfdf.@", "steeles80@gmail.ru"},
	EmailSearch{"ffff.@", "steeles80@gmail.ru"},
	EmailSearch{"FFF@.ff.f", "steeles80@gmail.ru"},
	EmailSearch{"steeles80@gmail.ru", "steeles80@gmail.ru"},
}

func TestEmailSearch(t *testing.T) {
	for i, test := range emailTest {
		if col := emailSearch(test.str); col != test.expected {
			t.Errorf("%d.Expected '%s', but got '%s'", i, test.expected, col)
		} else {
			fmt.Printf("%d---PASS---\n ", i)
		}
	}

}
