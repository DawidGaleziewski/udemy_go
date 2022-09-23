package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateAdminUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println("form name", name)

	ID, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
	}

	user := User{
		ID:           ID,
		Name:         name,
		Email:        email,
		Password:     password,
		CreationDate: time.Now(),
	}

	userData, validationErrors, err := user.Register()

	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(validationErrors) > 0 {
		errMsg := ""
		for _, value := range validationErrors {
			errMsg += value + "\n"
		}
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	redirectURL := fmt.Sprintf("/user/%v", userData.ID)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
