package main

import (
	"fmt"
	"regexp"
)

func currency(str string) string {
	rezult := ""
	s1, _ := regexp.MatchString(`EUR`, str)
	if s1 == true {
		re := regexp.MustCompile(`EUR[" "]\d{0,2}.\d{1,5}[.,]\d{1,3}`)
		res := re.FindAllString(str, -1)

		for _, v := range res {
			rezult += v
		}
	}
	s2, _ := regexp.MatchString(`$`, str)
	if s2 == true {
		re := regexp.MustCompile(`^\$\d{0,3}[.,]\d{1,3}`)
		res := re.FindAllString(str, -1)
		for _, v := range res {
			rezult += v
		}
	}
	fmt.Println(rezult)
	return rezult
}
func main() {
	currency("EUR 1.203.09")
	currency("$30.65")
	currency("$30.65 ($13.74 / Pound)")
	currency("$1,000,010.46")
	currency("+&nbsp;EUR 18,00&nbsp;di&nbsp;spedizione")
}
