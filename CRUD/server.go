package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	id   = "id"
	body = "body"
)

func main() {
	http.HandleFunc("/list", list)
	http.HandleFunc("/edit", edit)
	fmt.Println("start to run server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("Server is running...")
	}
	fmt.Println("finish application")
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
}

func edit(w http.ResponseWriter, r *http.Request) {
	// id, body
	fmt.Printf("%#v", r.URL.Query())
	values := r.URL.Query() // map[string][]string
	fmt.Println(values)
	if _, exist := values[id]; !exist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request"))
		return
	}
	if _, exist := values[body]; !exist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request"))
		return
	}
	// точно знаем что параметры тут есть
	// values[body] это слайс
	// values[id] это слайс
	if len(values[id]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request, id is empty"))
		return
	}
	if len(values[body]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request, body is empty"))
		return
	}
	// точно знаем что есть какие то данные в id и в body.
	_, err := strconv.Atoi(values[id][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request, id is not a number"))

		return

	}

	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error when opening file: %s", err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	count := 0
	str := ""
	for fileScanner.Scan() {
		number, _ := strconv.Atoi(values[id][0])
		if count == number {
			fmt.Println("tt")
			str = str + ((values[body][0]) + "\n")
		} else {
			str = str + (fileScanner.Text() + "\n")
		}

		count++
	}
	data := []byte(str)
	err = ioutil.WriteFile("test.txt", data, 0777) // запись измененых данных в фаил
	if err != nil {
		// Если произошла ошибка выводим ее в консоль
		fmt.Println(err)
	}
	w.Write([]byte("row is changed"))
	w.WriteHeader(http.StatusOK)
}
