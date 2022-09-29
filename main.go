package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"math/rand"
	"path/filepath"
	"encoding/hex"
	"os"
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
	secret := req.FormValue("text")
	randomFileName := randomFileName()

	f, err := os.Create(filepath.Join("secrets", randomFileName))
	if err != nil {
	    io.WriteString(w, "err 1!\n")
	    return
        }

        defer f.Close()

        _, err2 := f.WriteString(fmt.Sprintf("%s", secret))
	if err2 != nil {
	    io.WriteString(w, "err 2!\n")
	    return
        }
	
	io.WriteString(w, fmt.Sprintf("<html><body>Share secret with this ephemeral link (one-use): <a href=\"secret/%s\">%s</a>.</body></html>", randomFileName, randomFileName))
}

func secretShow(w http.ResponseWriter, req *http.Request) {
	slug := getField(req, 0)
	io.WriteString(w, slug)
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
    fields := r.Context().Value(ctxKey{}).([]string)
    return fields[index]
}

func randomFileName() string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return filepath.Join(hex.EncodeToString(randBytes))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", submit)
	http.HandleFunc("/secret/:slug", secretShow)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
