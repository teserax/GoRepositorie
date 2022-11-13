package main

import (
	"fmt"
	"sync"
)

func firstTeam(i int, ch chan int, wg *sync.WaitGroup, f func(i int) int) {
	defer wg.Done()
	ch <- f(i)
}
func sum(i int) int {
	return i + i
}

func main() {
	wg1 := new(sync.WaitGroup)
	do := true
	for do {

		var col [30]struct{}
		ch1 := make(chan int)

		for i := range col {
			wg1.Add(1)
			go firstTeam(i, ch1, wg1, sum)

		}
		go func() {
			defer close(ch1)
			wg1.Wait()
			do = false
		}()
		for v := range ch1 {
			fmt.Println(v)
		}

	}
}
