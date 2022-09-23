package session

import (
	dbutil "admin_panel/db_util"
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

func (session Session) Create() (Session, error) {
	var err error
	db, err := dbutil.Config.OpenConnection()
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
	);`, session.ID, session.UserID, dbutil.Config.FormatTime(session.CreationDate), dbutil.Config.FormatTime(session.LastActivityDate))

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

func (session Session) FindBy(query map[string]string) (sessionResults []Session, err error) {
	var dbSessionRecord Session
	db, err := dbutil.Config.OpenConnection()
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

	rows, err := db.Query(sqlQueryString)

	if err != nil {
		log.Println(err)
		return sessionResults, err
	}

	for rows.Next() {
		var creationTimeRaw []uint8
		var lastActivityTimeRaw []uint8
		err = rows.Scan(&dbSessionRecord.ID, &dbSessionRecord.UserID, &creationTimeRaw, &lastActivityTimeRaw)

		dbSessionRecord.CreationDate, err = time.Parse("2006-01-02 03:04:05", string(creationTimeRaw))
		if err != nil {
			log.Println("parsing creationDate", err)
		}

		dbSessionRecord.LastActivityDate, err = time.Parse("2006-01-02 03:04:05", string(lastActivityTimeRaw))
		if err != nil {
			log.Println("parsing LastActivityDate", err)
		}

		sessionResults = append(sessionResults, dbSessionRecord)
	}

	return sessionResults, err
}

func (session Session) Delete() (dbSessionRecord Session, err error) {
	db, err := dbutil.Config.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	fmt.Println("session recived: ", session)
	sessionDeleteStatment := fmt.Sprintf(`DELETE FROM gosql.sessions WHERE id="%v"`, session.ID)
	stmt, err := db.Prepare(sessionDeleteStatment)
	if err != nil {
		log.Println("error prepering statment", err)
		return session, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("error executing statment: ", sessionDeleteStatment, err)
		return session, err
	}

	return session, err
}
