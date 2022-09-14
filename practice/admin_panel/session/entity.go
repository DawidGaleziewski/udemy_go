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

func (session Session) Create(db *sql.DB) (dbSessionRecord Session, err error) {
	db, err = sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	if err != nil {
		log.Println(err)
		return dbSessionRecord, err
	}
	defer db.Close()

	sessionInsert := fmt.Sprintf(`
	INSERT INTO gosql.sessions (id, user_id, creation_date, last_activity)
	VALUES (
		'%v',
		'%v', 
		'%v',
		'%v'
	);`, dbSessionRecord.ID, dbSessionRecord.UserID, dbSessionRecord.CreationDate.UTC().Format("2006-01-02 03:04:05"), dbSessionRecord.LastActivityDate.UTC().Format("2006-01-02 03:04:05"))

	fmt.Println("executing query ", sessionInsert)

	stmt, err := db.Prepare(sessionInsert)
	if err != nil {
		log.Println(err)
		return dbSessionRecord, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return dbSessionRecord, err
	}
	return dbSessionRecord, err
}
