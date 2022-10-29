package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Label struct {
	Task struct {
		LabelID   string `json:"label_id"`
		ReleaseID string `json:"release_id"`
	} `json:"task"`
	Attributes struct {
		Country  string `json:"Country"`
		Genre    string `json:"Genre"`
		Label    string `json:"Label"`
		Released string `json:"Released"`
		Style    string `json:"Style"`
	} `json:"attributes"`
}

func main() {

	file, err := os.Open("parsed_releases.json") //открытие фала

	if err != nil {
		log.Fatalf("Error when opening file: %s", err) //проверка на ошибку открытия
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file) //чтение построчно
	bd := [][]byte{}
	for fileScanner.Scan() {

		bd = append(bd, []byte(strings.Trim(fileScanner.Text(), "\n")))
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err) //проверка на ошибку чтения
	}

	taskFile := createFile("task.txt")             //создание фаила
	attributesFile := createFile("attributes.txt") //создание фаила
	body := Label{}

	release := map[string]int{}
	for i := 1; i < len(bd); i++ {
		if err = json.Unmarshal(bd[i], &body); err != nil { //распарсивание JSON
			log.Fatal(err)
		}

		if _, exist := release[string(body.Task.LabelID)]; exist {
			release[string(body.Task.LabelID)]++
		}
		release[string(body.Task.LabelID)]++
		if body.Attributes.Country == "" || body.Attributes.Genre == "" || body.Attributes.Label == "" || body.Attributes.Released == "" || body.Attributes.Style == "" { //проверка на пустые значения
			continue
		} else {
			attributesFile.WriteString(fmt.Sprintf("Attributes: %3s\n", body.Attributes))
		}

	}
	type sortMap struct {
		key   string
		value int
	}
	sortReleas := []sortMap{}
	for rel, val := range release { //сортировка task
		sortReleas = append(sortReleas, sortMap{rel, val})
	}
	sort.Slice(sortReleas, func(i, j int) bool {
		return sortReleas[i].value > sortReleas[j].value
	})
	for _, val := range sortReleas {
		taskFile.WriteString(fmt.Sprintf("LabelId:%5s  ReleaseID:%5s\n", string(val.key), strconv.Itoa(val.value)))
	}

}
func createFile(nameFile string) os.File { //функция создания фаилов
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	return *file

} //THE END