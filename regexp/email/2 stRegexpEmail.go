package main

import (
	"regexp"
)

func emailSearch(s string) string {

	re, _ := regexp.Compile(`[a-z0-9\.]+@[a-z]+\.[a-z]{2,5}`)
	res := re.FindAllString(s, -1)
	for _, v := range res {

		return v
	}

	return ""
}
func main() {

}
