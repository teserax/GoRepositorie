package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	s := []string{}
	for fileScanner.Scan() {
		s = append(s, strings.Split(fileScanner.Text(), "\n")...)

	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	defer file.Close()
	number, err := strconv.Atoi(values[id][0])
	s[number] = values[body][0]
	fmt.Println(s)

	f, err := os.OpenFile("test.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	data := []byte{}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		data = append(data, []byte(s[i]+"\n")...)

	}
	err = ioutil.WriteFile("test.txt", data, 0777)
	if err != nil {
		// Если произошла ошибка выводим ее в консоль
		fmt.Println(err)
	}
	// найти нудЖную строку
	// перезаписать ее
	// сохранить файл
	// вернуть ответ что операция закончена успешно
	w.Write([]byte("row is changed"))

	w.WriteHeader(http.StatusOK)
}
