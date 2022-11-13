package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var channel = make(chan string)

func Wait() {
	wg.Wait()
	close(channel)
}

var ch = make([]chan int, 6, 6)

func main() {
	words := []string{
		"h", "e", "l", "l", "o",
	}
	for n := range ch {
		ch[n] = make(chan int)
	}
	go start()
	for i, v := range words {
		wg.Add(2)
		go takeWords(v, i)
		go PrintValue()

	}
	n := <-ch[5]
	n = n
	wg.Wait()

}
func start() {
	ch[0] <- 0
}

func takeWords(s string, i int) {
	defer wg.Done()
	val := <-ch[i]
	channel <- s

	ch[i+1] <- val

}
func PrintValue() {
	defer wg.Done()
	fmt.Print(<-channel)
}
