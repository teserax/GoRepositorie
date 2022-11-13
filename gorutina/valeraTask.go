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
		close(ch)
		interval.Stop()
		fmt.Println("timeout")
	case <-interval.C:
		coler := make([]struct{}, n)

		var wg1 = new(sync.WaitGroup)

		for range coler {
			wg1.Add(1)
			go func() {

				gn := rand.Intn(10) + 1
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

func numberProcessing(ch chan int) {
	coler := [10]struct{}{}

	var wg = new(sync.WaitGroup)
	for v := range ch {
		for range coler {
			wg.Add(1)
			go func() {
				fmt.Println(v)
				wg.Done()
			}()
		}
		go func() {
			wg.Wait()

		}()
	}

}

func main() {
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	ch2 := numbers(10, 1, 10, ch)
	numberProcessing(ch2)
}
