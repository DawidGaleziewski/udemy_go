package session

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	CreationDate     time.Time
	LastActivityDate time.Time
}

func (session Session) Create(db *sql.DB) (Session, error) {
	var err error
	db, err = sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	if err != nil {
		log.Println(err)
		return session, err
	}
	defer db.Close()
	// to do Verify that no more seesions then one exists for single user. If it does, destroy the last session
	sessionInsert := fmt.Sprintf(`
	INSERT INTO gosql.sessions (id, user_id, creation_date, last_activity)
	VALUES (
		'%v',
		'%v', 
		'%v',
		'%v'
	);`, session.ID, session.UserID, session.CreationDate.UTC().Format("2006-01-02 03:04:05"), session.LastActivityDate.UTC().Format("2006-01-02 03:04:05"))

	fmt.Println("executing query ", sessionInsert)

	stmt, err := db.Prepare(sessionInsert)
	if err != nil {
		log.Println(err)
		return session, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return session, err
	}
	return session, err
}
