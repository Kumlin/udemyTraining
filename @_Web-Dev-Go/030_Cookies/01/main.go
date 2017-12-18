package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// http.HandleFunc("/", set)
	// http.HandleFunc("/read", read)
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9190", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprintln(w, cookie)
}

// func set(w http.ResponseWriter, r *http.Request) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:  "my-cookie",
// 		Value: "some value",
// 	})
// 	fmt.Fprintln(w, "COOKIE WRITTEN - CHK BROWSER")
// 	fmt.Fprintln(w, "devtool / application / cookies")
// }
//
// func read(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("my-cookie")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
//
// 	fmt.Fprintln(w, "YOUR COOKIE:", c)
// }
