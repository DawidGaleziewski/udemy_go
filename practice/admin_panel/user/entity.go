package user

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	Password     string
	CreationDate time.Time
}

func (user User) validate() (validationErrors []string, isValid bool, err error) {
	const MAX_USER_LEN = 100
	const MIN_PASSWORD_LEN = 8
	//const MAX_PASSWORD_LEN = 1000

	if len(user.Name) == 0 {
		validationErrors = append(validationErrors, "user.Name can't be empty")
	}

	if len(user.Name) > MAX_USER_LEN {
		validationErrors = append(validationErrors, fmt.Sprint("user.Name can't be longer then %v", MAX_USER_LEN))
	}

	if len(user.Email) == 0 {
		validationErrors = append(validationErrors, "user.Email can't be empty")
	}

	match, err := regexp.MatchString(".*@.*\\..*", user.Email)
	if err != nil {
		log.Println("error parsing email regex")
		return validationErrors, isValid, err
	}

	if !match || err != nil {
		validationErrors = append(validationErrors, "user.Email is not a valid email")
	}

	if len(user.Password) == 0 {
		validationErrors = append(validationErrors, "user.Password can't be empty")
	}

	if len(user.Password) == 0 {
		validationErrors = append(validationErrors, "user.Password can't be empty")
	}

	if len(user.Password) < MIN_PASSWORD_LEN {
		validationErrors = append(validationErrors, fmt.Sprintf("user.Password must be at least %v characters long", MIN_PASSWORD_LEN))
	}

	isValid = len(validationErrors) == 0
	return validationErrors, isValid, err
}

func (user User) Register(db *sql.DB) (dbUserRecord User, validationErrors []string, err error) {
	var isUserDataValid bool

	validationErrors, isUserDataValid, err = user.validate()

	if err != nil {
		log.Println(err)
		return dbUserRecord, validationErrors, err
	}

	if !isUserDataValid {
		log.Println(err)
		return dbUserRecord, validationErrors, err
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
		return dbUserRecord, validationErrors, err
	}

	fmt.Println("before creating user")
	dbUserRecord = User{
		ID:           user.ID,
		Name:         user.Name,
		Password:     string(encryptedPassword),
		Email:        user.Email,
		CreationDate: user.CreationDate,
	}

	db, err = sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	userInsert := fmt.Sprintf(`
	INSERT INTO gosql.admin_users (id, name, email, password, creation_date)
	VALUES (
		'%v',
		'%v', 
		'%v',
		'%v',
		'%v'
	);`, dbUserRecord.ID, dbUserRecord.Name, dbUserRecord.Email, dbUserRecord.Password, dbUserRecord.CreationDate.UTC().Format("2006-01-02 03:04:05"))

	fmt.Println("executing query ", userInsert)

	stmt, err := db.Prepare(userInsert)
	if err != nil {
		log.Println(err)
		return dbUserRecord, validationErrors, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return dbUserRecord, validationErrors, err
	}

	return dbUserRecord, validationErrors, err
}
