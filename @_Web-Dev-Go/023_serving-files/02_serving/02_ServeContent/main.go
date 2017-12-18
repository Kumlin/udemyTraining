package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", pic)
	http.HandleFunc("/pic.png", picFull)
	http.ListenAndServe(":6565", nil)
}

func pic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src=/pic.png>`)
}

func picFull(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("pic.png")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
