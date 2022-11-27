package main

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	fmt.Println(result)
	result += " and "
	result += phrases[len(phrases)-1]
	return result
}
func main() {
	s := []string{"apple", "orange", "pear", "banana"}
	result := JoinWithCommas(s)

	fmt.Println(result)
}
