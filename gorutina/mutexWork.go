package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var do bool

func main() {
	var wg = new(sync.WaitGroup)
	ch := make(chan bool) // канал
	numCh := make(chan int)
	var mutex sync.Mutex // определяем мьютекс
	do = true

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			num := rand.Intn(3) + 1
			fmt.Println(num)
			work(num, ch, wg, &mutex)
			numCh <- num
		}()
		if <-ch == false {
			fmt.Println("we close")
			do = false
		}
	}

	go func() {
		close(numCh)
		wg.Wait()
	}()
	for i := range numCh {
		fmt.Println(i)
	}
	fmt.Println("The End")
}

func work(number int, ch chan bool, wg *sync.WaitGroup, mutex *sync.Mutex) {
	wg.Done()
	mutex.Lock() // блокируем доступ к переменной counter

	mutex.Unlock() // деблокируем доступ
	ch <- true
}
