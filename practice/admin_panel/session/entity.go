package session

import (
	dbutil "admin_panel/db_util"
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

func (session Session) FindBy(query dbutil.Query) (sessionResults []Session, err error) {
	var dbSessionRecord Session
	db, err := dbutil.Config.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query(query.Select("sessions"))

	if err != nil {
		log.Println(err)
		return sessionResults, err
	}

	for rows.Next() {
		var creationTimeRaw []uint8
		var lastActivityTimeRaw []uint8
		err = rows.Scan(&dbSessionRecord.ID, &dbSessionRecord.UserID, &creationTimeRaw, &lastActivityTimeRaw)

		dbSessionRecord.CreationDate, err = time.Parse(dbutil.Config.TimeLayout, string(creationTimeRaw))
		if err != nil {
			log.Println("parsing creationDate", err)
		}

		dbSessionRecord.LastActivityDate, err = time.Parse(dbutil.Config.TimeLayout, string(lastActivityTimeRaw))
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

	query := dbutil.Query{
		"id": session.ID.String(),
	}
	sessionDeleteStatment := query.Delete("sessions")

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
