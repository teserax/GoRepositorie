package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := []int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 9; i++ {
		num := rand.Intn(20) + 1
		s = append(s, num)
	}
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] > s[j] {
				t := s[i]
				s[i] = s[j]
				s[j] = t
			}
		}
	}
	fmt.Println(s)
	desiredNumber := rand.Intn(20) + 1
	fmt.Println("searh number:=", desiredNumber)
	var min, max int
	min = 0
	max = len(s) - 1

	for min <= max {
		middle := (min + max) / 2

		if desiredNumber == s[middle] {
			fmt.Println("Number found", s[middle])
			return
		} else if desiredNumber < s[middle] {
			max = middle - 1
		} else {
			min = middle + 1
		}

	}
	if min == len(s) || s[min] != desiredNumber {
		fmt.Println("No found")
	}
}
