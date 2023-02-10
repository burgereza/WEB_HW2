package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "404303202101", "HW2")

// create new user in database
func InsertUser(username string, password string, phoneNumber string) {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	_, err = db.Exec("INSERT INTO users(username, password ,phonenumber) VALUES($1,$2,$3)", username, password, phoneNumber)

	if err != nil {
		panic(err)
	}
}

// see if password is correct
func CheckUserInfo(username string, password string) bool {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	var db_pass string
	// err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&db_pass)
	err = db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&db_pass)

	if err != nil {
		return false
	}

	if db_pass != password {
		return false
	}

	return true
}

func GetCities() []string {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	s := make([]string, 3)

	rows, err := db.Query("SELECT city_name FROM city WHERE country_name=iran")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var city_name string
		err := rows.Scan(&city_name)
		if err != nil {
			panic(err)
		}
		s = append(s, city_name)
	}

	return s
}

// for unique username
func CheckPhoneNumber(phoneNumber string) bool {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)
	var x string
	err = db.QueryRow("SELECT phonenumber FROM users WHERE phonenumber=$1", phoneNumber).Scan(&x)

	if err != nil {
		return true
	}

	return false
}

// for unique email
func CheckEmailAvailability(username string) bool {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	var x string
	err = db.QueryRow("SELECT email FROM users WHERE username=?", username).Scan(&x)

	if err != nil {
		return true
	}

	return false
}

func DB_conn() (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "404303202101", "HW2")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}
