package main

// we can distribute same code in diffrent files and keep the same package name. As long as those are in the same folder. But we need to run it with "go run .". ALso may beed to re-init go mod init
import "net/http"

func IsLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	sessionID := sessionDB[cookie.Value] // we have access to global  variables as well
	_, ok := userDB[sessionID]
	return ok
}