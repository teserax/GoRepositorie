package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Personat struct {
	Id      string
	FIO     string
	Address string
	Product Products
}

type Products struct {
	Goods    []string
	Quantity int
	Price    string
}

func main() {
	file, err := os.Open("text/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var persons []Personat
	for scanner.Scan() {
		//   fmt.Println(scanner.Text())
		var person Personat
		err = json.Unmarshal(scanner.Bytes(), &person)
		persons = append(persons, person)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
