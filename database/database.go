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
func InsertUser(firstname, lastname, password, phoneNumber string) {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	_, err = db.Exec("INSERT INTO users(firstname ,lastname , password ,phonenumber) VALUES($1,$2,$3,$4)", firstname, lastname, password, phoneNumber)

	if err != nil {
		panic(err)
	}
	db.Close()
}

// see if password is correct
func CheckUserInfo(phonenumber string, password string) bool {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	var db_pass string
	// err := db.QueryRow("SELECT password FROM users WHERE phonenumber=?", phonenumber).Scan(&db_pass)
	err = db.QueryRow("SELECT password FROM users WHERE phonenumber=$1", phonenumber).Scan(&db_pass)

	if err != nil {
		return false
	}

	if db_pass != password {
		return false
	}
	db.Close()
	return true
}

func GetCities() []string {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	// s := make([]string, 3)
	var s []string

	rows, err := db.Query("SELECT city_name FROM airport")

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
	db.Close()
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
	db.Close()
	return false
}

// for unique email
func CheckEmailAvailability(username string) bool {
	// var db *sql.DB
	// db = DB_conn()

	db, err := sql.Open("postgres", psqlInfo)

	var x string
	err = db.QueryRow("SELECT email FROM users WHERE username=$1", username).Scan(&x)

	if err != nil {
		return true
	}
	db.Close()
	return false
}

func DB_conn() {

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "404303202101", "HW2")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//return db
}
