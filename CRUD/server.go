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
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/get", get)
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
func get(w http.ResponseWriter, r *http.Request) { //получаем содержимое страницы....тут "фаила"

	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error when opening file: %s", err)
	}
	defer file.Close()
	var str string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		str += fileScanner.Text() + "\n"
	}
	data := []byte(str)
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func delete(w http.ResponseWriter, r *http.Request) { //удаляем содержимое по ID на страницы....тут "фаила"
	fmt.Printf("%#v", r.URL.Query())
	values := r.URL.Query() // map[string][]string
	fmt.Println(values)
	if _, exist := values[id]; !exist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request"))
		return
	}

	if len(values[id]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect request, id is empty"))
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
			str = str + "\n"
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

func edit(w http.ResponseWriter, r *http.Request) { //добавляем строку по-указанному ID согласно запроса body на страницы....тут "фаила"
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
