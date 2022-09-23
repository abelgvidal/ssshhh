package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	var tmplt *template.Template
	tmplt, err := template.ParseFiles("form.html")
	if err != nil {
		io.WriteString(w, "err!\n")
		return
	}
	tmplt.Execute(w, "")
}

func submit(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf("%s", reqBody))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", submit)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
