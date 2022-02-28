package main

import (
	"fmt"
	"io/ioutil"
)

type Persona struct {
	Id      string
	Fio     string
	Address string
	Product Product
}

//type Adress struct {
//	City   string
//	street string
//}
type Product struct {
	Name     string
	Quantity int
	Price    string
}

func main() {

	fContent, err := ioutil.ReadFile("1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
