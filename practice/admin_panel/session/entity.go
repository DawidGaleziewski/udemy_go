package session

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

func (session Session) FindBy(db *sql.DB, query map[string]string) (sessionResults []Session, err error) {
	var dbSessionRecord Session
	db, err = sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	sqlQueryString := `SELECT * from gosql.sessions`
	var whereConditionsSlice []string
	for rowName, condition := range query {
		whereConditionsSlice = append(whereConditionsSlice, fmt.Sprintf("%v=\"%v\"", rowName, condition))

	}

	whereConditionsQuery := ` WHERE `
	if len(whereConditionsSlice) > 0 {
		whereConditionsQuery += strings.Join(whereConditionsSlice, "AND ")
	}

	if len(whereConditionsQuery) > 0 {
		sqlQueryString += whereConditionsQuery
	}

	fmt.Println("using query", sqlQueryString)
	rows, err := db.Query(sqlQueryString)

	if err != nil {
		log.Println(err)
		return sessionResults, err
	}

	for rows.Next() {
		var creationTimeRaw []uint8
		var lastActivityTimeRaw []uint8
		err = rows.Scan(&dbSessionRecord.ID, &dbSessionRecord.UserID, &dbSessionRecord.CreationDate, &creationTimeRaw, &lastActivityTimeRaw)

		parsedCreationDate, err := time.Parse("2006-01-02 03:04:05", string(creationTimeRaw))
		if err != nil {
			log.Println(err)
			return sessionResults, err
		}
		dbSessionRecord.CreationDate = parsedCreationDate

		parsedLastActivityTimeRaw, err := time.Parse("2006-01-02 03:04:05", string(creationTimeRaw))
		if err != nil {
			log.Println(err)
			return sessionResults, err
		}
		dbSessionRecord.LastActivityDate = parsedLastActivityTimeRaw

		sessionResults = append(sessionResults, dbSessionRecord)
	}

	return sessionResults, err
}

func (session Session) Delete(db *sql.DB) (dbSessionRecord Session, err error) {
	sessionDeleteStatment := fmt.Sprintf(`DELETE FROM gosql.sessions WHERE id="%v"`, session.ID)

	stmt, err := db.Prepare(sessionDeleteStatment)
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
