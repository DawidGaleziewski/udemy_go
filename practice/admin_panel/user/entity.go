package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
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

type QueryUser map[string]string

func (user User) validate(db *sql.DB) (validationErrors []string, isValid bool, err error) {
	const MAX_USER_LEN = 100
	const MIN_PASSWORD_LEN = 8
	//const MAX_PASSWORD_LEN = 1000

	rows, err := db.Query(`SELECT email from gosql.admin_users`)

	var cellValue string
	for rows.Next() {
		err = rows.Scan(&cellValue) // stoped here
	}

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
	db, err = sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var isUserDataValid bool

	validationErrors, isUserDataValid, err = user.validate(db)

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

func (user User) FindBy(db *sql.DB, query QueryUser) (usersResult []User, err error) {
	var dbUserRecord User
	db, err = sql.Open("mysql", "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	sqlQueryString := `SELECT * from gosql.admin_users`
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
		return usersResult, err
	}

	for rows.Next() {
		var creationTimeRaw []uint8
		err = rows.Scan(&dbUserRecord.ID, &dbUserRecord.Name, &dbUserRecord.Email, &dbUserRecord.Password, &creationTimeRaw)
		parsedCreationDate, err := time.Parse("2006-01-02 03:04:05", string(creationTimeRaw))
		dbUserRecord.CreationDate = parsedCreationDate
		if err != nil {
			log.Println(err)
			return usersResult, err
		}
		fmt.Println("record", dbUserRecord, "time raw")
		if err != nil {
			log.Println(err)
			return usersResult, err
		}

		usersResult = append(usersResult, dbUserRecord)
	}

	return usersResult, err
}

func (user User) VerifyCredentials(db *sql.DB, email string, password string) (verifiedUser User, isVerified bool, err error) {
	usersFound, err := user.FindBy(db, map[string]string{
		"email": email,
	})

	if err != nil {
		log.Println(err)
		return user, false, err
	}

	if len(usersFound) == 0 {
		return user, false, err
	}

	var verifiedUsers []User
	for _, dbUserRecord := range usersFound {
		isValidPassword := bcrypt.CompareHashAndPassword([]byte(dbUserRecord.Password), []byte(password)) == nil
		if dbUserRecord.Email == email && isValidPassword {
			verifiedUsers = append(verifiedUsers, dbUserRecord)
		}
	}

	if len(verifiedUsers) > 2 {
		return user, false, errors.New("db should newer return more then two users verified by one credentials")
	}

	if len(verifiedUsers) == 0 {
		return user, false, err
	}

	return verifiedUser, true, err
}
