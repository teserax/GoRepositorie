package main

import (
	"fmt"
	"os"
)

func main() {
	//	text := "Hellhhhhhhhhhhhho Gold!"
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	str := "rtteeeeeeeeeeeeeeeeeeeeeeeeeeeettt"
	file.WriteString(str + "\n")
	file.WriteString(str)

	fmt.Println("Done.")
}
