package main

import (
	"fmt"
	"net/http"
)

var Task1 = "Finish internship work"
var Task2 = "Practice backend engineering in Golang and kubernetes"
var Task3 = "dive deeper in companies code base"
var taskItems = []string{Task1, Task2, Task3}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}
