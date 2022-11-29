package main

import (
	"fmt"
	"regexp"
)

func priceTranslation(s string) string {

	re, _ := regexp.Compile(`EUR|\$\d*.|\s\d*.\d*`)
	res := re.FindAllString(s, -1)
	var rezult string
	for _, v := range res {
		rezult += v
	}
	fmt.Printf(rezult)
	return rezult
}
func main() {

}
