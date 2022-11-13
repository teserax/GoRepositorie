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

var wg sync.WaitGroup

//   - первый этап: N1 количество горутин генерируют случайные числа с интервалом N2 секунд
//     в каком то диапазоне [a,b] и отправляет в канал 1
func randomNumberGeneration(numberOfGourutin, interval int) <-chan int {
	out := make(chan int, 1)
	for i := 0; i < numberOfGourutin; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(interval))
			num := rand.Intn(10) + 1
			fmt.Println(num)
			out <- num
			wg.Done()
		}()

		select {
		case n := <-out:
			if n == 10 {
				fmt.Println("Ten")
			}
			close(out)
			return out
		default:
			break
		}
	}
	return out

}
func main() {
	rand.Seed(time.Now().UnixNano())

	result := randomNumberGeneration(10, 1)
	result = result
	//for n := range result {
	//	fmt.Print(n)
	//}

	wg.Wait()

}
