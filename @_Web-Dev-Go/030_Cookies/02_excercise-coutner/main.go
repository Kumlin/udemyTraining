package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", enter)
	http.HandleFunc("/number", number)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9893", nil)
}

func enter(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("counter")
	if c.Value == "" {
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: "0",
		})
		return
	}
	currentCount, _ := strconv.Atoi(c.Value)
	newCount := currentCount + 1
	c.Value = strconv.Itoa(newCount)
	http.SetCookie(w, c)
	fmt.Fprintln(w, c)
}

func number(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
