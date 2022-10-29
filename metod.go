package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Persona struct {
	Id      string
	FIO     string
	Address string
}

func (p *Persona) addFio(FIO string) {

	p.FIO = FIO

}
func (p *Persona) addAddress(Address string) {

	p.Address = Address

}

func main() {
	buyer, err := readFile("text/1.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	fmt.Println(buyer[1].Id)

}

//read file and unmarshal json
func readFile(a string) ([]Persona, error) {
	// open the file
	file, err := os.Open(a)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var persons []Persona
	for scanner.Scan() {
		//   fmt.Println(scanner.Text())
		var person Persona
		err = json.Unmarshal(scanner.Bytes(), &person)
		persons = append(persons, person)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return persons, nil
}
