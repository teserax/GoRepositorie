// - первый этап: N1 количество горутин генерируют случайные числа с интервалом N2 секунд
//    в каком то диапазоне [a,b] и отправляет в канал 1
// - второй этап: M1 гороутин берут эти числа из канала 1
//    и делают какую то обработку их, например находят все делители этих чисел и отправляют json структуру (число и его делители) в канал 2
// - третий этап: K1 гороутин получают json и выделяют
//    максимальное число переданной на 1вом этапе.
//--программа заканчивает работу после В секунд или если сгеренируется на первом этапе максимальное число с интервала --

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func numbers(n int, inter, workTime time.Duration, ch chan int) chan int {
	interval := time.NewTicker(time.Second * inter)
	timer := time.After(time.Second * workTime)

	select {
	case <-timer:
		fmt.Println("timeout")
		close(ch)
		interval.Stop()

	case <-interval.C:
		coler := make([]struct{}, n)

		var wg1 = new(sync.WaitGroup)
		maxNumberGenerat := 10
		for range coler {
			wg1.Add(1)
			go func() {
				gn := rand.Intn(maxNumberGenerat) + 1
				ch <- gn
				wg1.Done()
			}()
		}
		go func() {
			defer close(ch)
			wg1.Wait()
		}()
	}
	return ch
}
func mult(ch chan int) {
	for range ch {

	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	rand.Seed(time.Now().UnixNano())
	numbers(10, 1, 1, ch)
	var wg = new(sync.WaitGroup)

	for v := range ch {
		if v == 10 {
			fmt.Println("maximum number generated exit")
			return
		} else {
			ch2 <- v
			go mult(ch2)
		}

	}
	go func() {
		defer close(ch)
		wg.Wait()

	}()

}
