package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy dogity")
	case "/cat":
		io.WriteString(res, "kitty kitty katty")
	default:
		io.WriteString(res, "no case")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":4949", d)
}
