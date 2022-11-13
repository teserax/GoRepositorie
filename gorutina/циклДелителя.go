package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	m := make(map[int][]int)
	s := []int{}
	for i := 0; i < 10; i++ {
		num := rand.Intn(9) + 2
		s = append(s, num)
	}
	fmt.Println(s)
	for _, v := range s {
		//m[v] = append(m[v])
		for i := 2; i < len(s); i++ {
			if v%i == 0 {
				if _, ok := m[v]; ok {
				} else {
					m[v] = append(m[v], i)
				}
			}
		}
	}
	fmt.Println(m)

}
