package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)

	var u user
	if username, ok := dbSessions[c.Value]; ok {
		u = dbUsers[username]
	}
	return u
}

func alreadyLoggedIn(request *http.Request) bool {
	cookie, err := request.Cookie("session")
	if err != nil {
		return false
	}
	username := dbSessions[cookie.Value]
	_, ok := dbUsers[username]
	return ok
}
