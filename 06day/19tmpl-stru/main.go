package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":9090", nil)
}

// 解析结构体
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		log.Fatal(err)
	}
	s := Student{
		Name:   "zhangsan",
		Age:    23,
		Gender: "nan",
	}
	t.Execute(w, s)
}

type Student struct {
	Name   string
	Age    int
	Gender string
}
