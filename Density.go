//Известны данные о массе (в кг) и объеме (в см3) 20 тел,
//изготовленных  из  различных  материалов.
//Определить  макси-мальную плотность материала
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Material struct {
	Masa    float64
	Volum   float64
	Densiti float64
}

func main() {
	products := []Material{} //make slice of products
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ { //in for rand data slice
		p := Material{
			Masa:  rand.Float64() * 100,
			Volum: rand.Float64() * 100,
		}
		products = append(products, p)
	}
	var maxDensiti Material
	for i := range products {
		products[i].Densiti = products[i].Masa / products[i].Volum
		if maxDensiti.Densiti < products[i].Densiti {
			maxDensiti = products[i]
		}
	}
	fmt.Println("----------------------------")
	fmt.Printf("Masa:%15.2f\nVolum:%14.2f\nMax Densiti:%8.2f\n", maxDensiti.Masa, maxDensiti.Volum, maxDensiti.Densiti)
	fmt.Println("----------------------------")
}
