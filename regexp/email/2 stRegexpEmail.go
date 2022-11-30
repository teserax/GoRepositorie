package main

import (
	"log"
	"regexp"
	"strings"
)

func emailSearch(s string) string {
	str := strings.Fields(s)
	for _, word := range str {
		email, err := regexp.MatchString(`[a-zA-Z0-9]@[a-z]`, word)
		if err != nil {
			log.Fatal(err)
		}
		if email == true {
			//fmt.Println(word)
			re, _ := regexp.Compile(`\w*\100[a-z]*[.][a-z]*`)
			res := re.FindAllString(word, -1)
			for _, v := range res {

				return v
			}

		}
	}
	return ""
}
func main() {
	//s := "dadadad steeles80@@@m34ail.ru ijsifsdfsd steeles80@mail.ru steeles80@m34ail..ru qeqeqw eqeqweq@sfsdf sadad42@3df@dfgdf.ouo stee/[les80@m34ail.ru"

}
