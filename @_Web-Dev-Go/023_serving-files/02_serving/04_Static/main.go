package main

import "net/http"

func main() {
	http.ListenAndServe(":8880", http.FileServer(http.Dir(".")))
}
