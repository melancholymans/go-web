package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	tf, err := template.ParseFiles("templates/hello1.html")
	if err != nil {
		tf, _ = template.New("index").Parse(("<html><body><h1>NO TEMP</h1></body></html>"))
	}

	hh := func(w http.ResponseWriter, rq *http.Request) {
		err = tf.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/hello", hh)
	http.ListenAndServe("", nil)
}
