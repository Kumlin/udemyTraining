package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", key)
	http.HandleFunc("/key.png", keyPic)
	http.ListenAndServe(":9970", nil)
}

func key(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
    <img src="/key.png">
    `)
}

func keyPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("key.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
