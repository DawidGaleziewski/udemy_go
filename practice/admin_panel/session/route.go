package session

import (
	"admin_panel/user"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	u := user.User{}

	verifiedUser, isAuthn, err := u.VerifyCredentials(email, password)

	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if !isAuthn {
		http.Error(w, "wrong credentials", http.StatusUnauthorized)
		return
	}

	id := uuid.New()
	newSession := Session{
		ID:               id,
		UserID:           verifiedUser.ID,
		CreationDate:     time.Now(),
		LastActivityDate: time.Now(),
	}

	sessionRecord, err := newSession.Create()

	if err != nil {
		log.Println(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionRecord.ID.String(),
		Path:  "/",
	})

	return
}
