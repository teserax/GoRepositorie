package main

import (
	"fmt"
	"time"
)

func message(ch chan string) {

	fmt.Println("Work")
	ch <- "Info"

	ch <- "Info"
	fmt.Println("Work")
}

func main() {
	ch := make(chan string)
	func() {
		go message(ch)
		select {
		case m := <-ch:
			fmt.Printf("Channel message %s\n", m)

		case <-time.After(1 * time.Second):
			fmt.Println("Time finish")
			fmt.Println("Channel close")
		}
	}()
	fmt.Println("Close bay")
}
