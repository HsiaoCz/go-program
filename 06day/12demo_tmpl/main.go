package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", HandleTemplate)

	http.Handle("/css/", http.FileServer(http.Dir("wwwrot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwrot")))
}

func LoadTemplates() *template.Template {
	result := template.New("templates")
	result, err := result.ParseGlob("templates/*.html")
	// 使用template.Must()这个函数处理错误
	template.Must(result, err)
	return result
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	template := LoadTemplates()
	filename := r.URL.Path[1:]
	t := template.Lookup(filename)
	if t != nil {
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
