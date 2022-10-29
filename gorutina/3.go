package main

import (
	"fmt"
	"sync"
)

var transfer_sum_ch chan int = make(chan int)
var main_sum_ch chan int = make(chan int)
var sum int = 0

var wg sync.WaitGroup

func Wait() {
	wg.Wait()
	close(transfer_sum_ch)
}

func transfer(v int) {
	defer wg.Done()
	fmt.Println("v-> ", v)
	transfer_sum_ch <- v
}
func Sum() {
	rez := 0
	i := 0
	for v := range transfer_sum_ch {
		if i == 0 {
			rez -= v
		} else {
			rez += v
		}
		i++
	}
	main_sum_ch <- rez
}

func main() {

	a := 1
	b := -1

	wg.Add(1)
	go transfer(a)
	wg.Add(1)
	go transfer(b)
	go Wait()
	go Sum()

	sum = <-main_sum_ch

	fmt.Println("sum-> ", sum)

}
