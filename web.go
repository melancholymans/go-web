package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	html := `<html>
	<body>
	<h1>HELLO</h1>
	<p>This is sample message</p>
	</body>
	</html>`
	tf, err := template.New("index").Parse(html)
	if err != nil {
		log.Fatal(err)
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
