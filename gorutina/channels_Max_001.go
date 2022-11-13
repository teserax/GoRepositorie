package main

import (
	"fmt"
	"sync"
	"time"
)

func firstGroup(i int, ch chan int, wg *sync.WaitGroup, f func(int) int) {
	defer wg.Done()
	ch <- f(i)

}
func secondGroup(ch1, ch2 chan int, wg *sync.WaitGroup, f func(int) int) {
	defer wg.Done()
	for v := range ch1 {
		ch2 <- f(v)
	}
}
func inc(i int) int {
	return i + 1
}
func mul(i int) int {
	return i * 2
}
func main() {
	do := true
	step := 0
	t := time.NewTicker(time.Second * 1)
	tAfter := time.After(time.Second * 8)

	for do {
		select {
		case <-tAfter:
			t.Stop()
			do = false
			fmt.Println("time after done")
		default:
			<-t.C
			step++
			var col [30]struct{}
			ch1 := make(chan int)
			var wg1 = new(sync.WaitGroup)

			for i := range col {
				wg1.Add(1)
				go firstGroup(i, ch1, wg1, inc)

			}
			go func() {
				defer close(ch1)
				wg1.Wait()

			}()

			var col2 [10]struct{}
			ch2 := make(chan int)
			var wg2 = new(sync.WaitGroup)
			for range col2 {
				wg2.Add(1)
				go secondGroup(ch1, ch2, wg2, mul)
			}
			go func() {
				defer close(ch2)
				wg2.Wait()
			}()
			i := 0
			var slice []int
			for v := range ch2 {
				slice = append(slice, v)
				if i == 29 {
					if v == 50 {
						fmt.Println("done step ", step)
						do = false
						t.Stop()
						fmt.Println(slice)
					} else {
						//fmt.Println("one more time")
					}
				}
				//fmt.Println(v)
				i++
			}
			//fmt.Println(slice)
		}

	}
}
